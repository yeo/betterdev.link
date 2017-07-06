defmodule Betterdev.Community.Collection do
  use Ecto.Schema
  import Ecto.Changeset
  alias Betterdev.Community.Collection
  alias Betterdev.Accounts.User
  alias Betterdev.Community.Link


  schema "community_collections" do
    field :name, :string
    belongs_to :user, User

    many_to_many :links, Link, join_through: "community_collection_links", on_replace: :delete
    timestamps()
  end

  @doc false
  def changeset(%Collection{} = collection, attrs) do
    collection
    |> cast(attrs, [:name])
    |> validate_required([:name])
  end
end
