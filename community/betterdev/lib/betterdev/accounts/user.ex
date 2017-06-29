defmodule Betterdev.Accounts.User do
  use Ecto.Schema
  import Ecto.Changeset
  alias Betterdev.Accounts.User
  alias Betterdev.Community.Link
  alias Betterdev.Community.Collection

  schema "accounts_users" do
    field :email, :string
    field :name, :string

    field :jwt_sub, :string
    field :jwt_aud, :string

    has_many :links, Link
    has_many :collections, Collection

    timestamps()
  end

  @doc false
  def changeset(%User{} = user, attrs) do
    user
    |> cast(attrs, [:email, :name, :jwt_aud, :jwt_sub])
    |> validate_required([:email, :name])
  end
end
