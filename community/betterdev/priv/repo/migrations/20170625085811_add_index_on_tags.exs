defmodule Betterdev.Repo.Migrations.AddIndexOnTags do
  use Ecto.Migration

  def change do
    create index(:community_tags, [:title], unique: true)
    create index(:community_link_tags, [:link_id, :tag_id], unique: true)
  end
end
