defmodule Betterdev.Community do
  @moduledoc """
  The boundary for the Community system.
  """

  import Ecto.Query, warn: false
  alias Betterdev.Repo

  alias Betterdev.Community.Link

  @doc """
  Returns the list of links.

  ## Examples

      iex> list_links()
      [%Link{}, ...]

  """
  def list_links(params \\ %{}) do
    #Link |> Repo.all
    Link |> Repo.paginate(params)
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
    w = Scrape.website(uri)
    if w.title do
      #%{user: user, title: w.title || url, uri: url, description: w.description, picture: w.image || w.favicon, status: "published", } |> Repo.insert()
      %Link{user: user}
        |> Link.changeset(%{title: w.title, uri: uri, description: w.description, picture: w.image || w.favicon, status: "published"})
        |> Repo.insert()
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
