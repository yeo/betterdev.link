defmodule Betterdev.Repo.Migrations.CreateTags do
  use Ecto.Migration

  def change do
    create table(:community_tags) do
      add :title, :string
      add :type, :string

      timestamps()
    end

    create table(:community_link_tags, primary_key: false) do
      add :link_id, :integer
      add :tag_id, :integer

      timestamps()
    end
  end
end
