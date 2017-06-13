defmodule Betterdev.Repo.Migrations.CreateBetterdev.Community.Link do
  use Ecto.Migration

  def change do
    create table(:community_links) do
      add :title, :string
      add :uri, :string

      timestamps()
    end

  end
end
