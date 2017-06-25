defmodule Betterdev.Community.Tag do
  use Ecto.Schema
  import Ecto.Changeset
  alias Betterdev.Community.Tag
  alias Betterdev.Community.Link

  schema "community_tags" do
    field :title, :string
    field :type, :string

    many_to_many :links, Link, join_through: "community_link_tags"

    timestamps()
  end

  @doc false
  def changeset(%Tag{} = tag, attrs) do
    tag
    |> cast(attrs, [:title])
    |> validate_required([:title])
  end
end
