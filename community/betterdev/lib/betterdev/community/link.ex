defmodule Betterdev.Community.Link do
  use Ecto.Schema
  import Ecto.Changeset
  alias Betterdev.Community.Link


  schema "community_links" do
    field :title, :string
    field :uri, :string

    timestamps()
  end

  @doc false
  def changeset(%Link{} = link, attrs) do
    link
    |> cast(attrs, [:title, :uri])
    |> validate_required([:title, :uri])
  end
end
