defmodule Betterdev.Community do
  @moduledoc """
  The boundary for the Community system.
  """

  import Ecto.Query, warn: false
  alias Betterdev.Repo

  alias Betterdev.Community.Link
  alias Betterdev.Community.Tag
  import Algolia

  @doc """
  Returns the list of links.

  ## Examples

      iex> list_links()
      [%Link{}, ...]

  """
  def list_links(params \\ %{}) do
    #Link |> Repo.all
    link = from(p in Link, order_by: [desc: :id], preload: :tags)
    link |> Repo.paginate(params)
  end

  @doc """
  Gets a single link.

  Raises `Ecto.NoResultsError` if the Link does not exist.

  ## Examples

      iex> get_link!(123)
      %Link{}

      iex> get_link!(456)
      ** (Ecto.NoResultsError)

  """
  def get_link!(id), do: Repo.get!(Link, id)

  @doc """
  Creates a link.

  ## Examples

      iex> create_link(%{field: value})
      {:ok, %Link{}}

      iex> create_link(%{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
   ## %Link{user_id: 1, title: w.title || url, uri: url, description: w.description, picture: w.image || w.favicon, status: "published", } |> Repo.insert()
  def create_link(attrs \\ %{}) do
    %Link{}
    |> Link.changeset(attrs)
    |> Repo.insert()
  end

  @doc """
  Create a link with uri only. We will parse title, description from the link
  """
  def user_post_link(user, attrs) do
    # TODO Use task/job queue and return to client instantly via web socket
    uri = attrs["uri"]
    w = Scrape.article(uri)
    if w.title do
      #%{user: user, title: w.title || url, uri: url, description: w.description, picture: w.image || w.favicon, status: "published", } |> Repo.insert()
      {:ok, link} = %Link{user: user}
        |> Link.changeset(%{title: w.title, uri: uri, description: w.description, picture: w.image || w.favicon, status: "published"})
        |> Repo.insert()
      Task.start_link(fn -> post_process_link(link) end)
      {:ok, link}
    end
  end

  @doc """
  Post processing once a link is submited.

  We will:
   - index to algolia
   - process tag
  """
  def post_process_link(link) do
    w = Scrape.article(link.uri)
    r = %{ objectID: link.id, id: link.id, title: link.title, description: link.description, content: w.fulltext, }
    "community" |> save_object(r)

    tags = Betterdev.Helper.Classifier.extract(w.fulltext)
    # Insert tag
    tags = tags ++ (w.tags |> Enum.filter_map(&(&1[:accuracy] >= 0.7), &(&1[:name])))
    tags |> Enum.map(fn (title) ->
      t = retreive_tag(title)
      # http://blog.roundingpegs.com/an-example-of-many-to-many-associations-in-ecto-and-phoenix/
      # We need preload to preapre for changset below
      t = t |> Repo.preload(:links)
      link = link |> Repo.preload(:tags)
      try do
        changeset = Ecto.Changeset.change(link) |> Ecto.Changeset.put_assoc(:tags, [t])
        Repo.update!(changeset)
      rescue
        e in RuntimeError -> e
      end
    end)
  end

  def retreive_tag(t) do
    case Repo.get_by(Tag, title: t) do
      nil -> %Tag{title: t, type: "autogen"} |> Repo.insert!()
      tag -> tag
    end
  end

  @doc """
  Updates a link.

  ## Examples

      iex> update_link(link, %{field: new_value})
      {:ok, %Link{}}

      iex> update_link(link, %{field: bad_value})
      {:error, %Ecto.Changeset{}}

  """
  def update_link(%Link{} = link, attrs) do
    link
    |> Link.changeset(attrs)
    |> Repo.update()
  end

  @doc """
  Deletes a Link.

  ## Examples

      iex> delete_link(link)
      {:ok, %Link{}}

      iex> delete_link(link)
      {:error, %Ecto.Changeset{}}

  """
  def delete_link(%Link{} = link) do
    Repo.delete(link)
  end

  @doc """
  Returns an `%Ecto.Changeset{}` for tracking link changes.

  ## Examples

      iex> change_link(link)
      %Ecto.Changeset{source: %Link{}}

  """
  def change_link(%Link{} = link) do
    Link.changeset(link, %{})
  end
end
