defmodule Betterdev.Community.Collection do
  use Ecto.Schema
  import Ecto.Changeset
  alias Betterdev.Community.Collection
  alias Betterdev.Accounts.User


  schema "community_collections" do
    field :name, :string
    belongs_to :user, User
    timestamps()
  end

  @doc false
  def changeset(%Collection{} = collection, attrs) do
    collection
    |> cast(attrs, [:name])
    |> validate_required([:name])
  end
end
