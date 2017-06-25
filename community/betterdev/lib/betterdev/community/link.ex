defmodule Betterdev.Community.Link do
  use Ecto.Schema
  import Ecto.Changeset
  alias Betterdev.Community.Link
  alias Betterdev.Community.Tag
  alias Betterdev.Accounts.User


  schema "community_links" do
    field :title, :string
    field :uri, :string
		field :picture, :string
		field :status, :string
		field :description, :string

    belongs_to :user, User
    many_to_many :tags, Tag, join_through: "community_link_tags"
    timestamps()
  end

  @doc false
  def changeset(%Link{} = link, attrs) do
    link
    |> cast(attrs, [:title, :uri, :picture, :status, :description])
    |> validate_required([:title, :uri, :user])
  end
end
