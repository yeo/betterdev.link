defmodule Betterdev.Accounts.User do
  use Ecto.Schema
  import Ecto.Changeset
  alias Betterdev.Accounts.User
  alias Betterdev.Community.Link

  schema "accounts_users" do
    field :email, :string
    field :name, :string

    has_many :links, Link
    timestamps()
  end

  @doc false
  def changeset(%User{} = user, attrs) do
    user
    |> cast(attrs, [:email, :name])
    |> validate_required([:email, :name])
  end
end
