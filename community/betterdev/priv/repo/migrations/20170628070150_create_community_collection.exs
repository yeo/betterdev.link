defmodule Betterdev.Repo.Migrations.CreateBetterdev.Community.Collection do
  use Ecto.Migration

  def change do
    create table(:community_collections) do
      add :name, :string
      add :user_id, :integer

      timestamps()
    end

  end
end
