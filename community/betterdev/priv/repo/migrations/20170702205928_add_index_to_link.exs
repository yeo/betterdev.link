defmodule Betterdev.Repo.Migrations.AddIndexToLink do
  use Ecto.Migration

  def change do
    create index(:community_links, [:uri], unique: true)
  end
end
