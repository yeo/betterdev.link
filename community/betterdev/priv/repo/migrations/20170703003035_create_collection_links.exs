defmodule Betterdev.Repo.Migrations.CreateCollectionLinks do
  use Ecto.Migration

  def change do
    create table(:community_collection_links, primary_key: false) do
      add :link_id, :integer
      add :collection_id, :integer

      timestamps()
    end

    create index(:community_collection_links, [:link_id, :collection_id], unique: true)
  end
end
