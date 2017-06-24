defmodule Betterdev.Repo.Migrations.AddIndexToUsers do
  use Ecto.Migration

  def change do
    create index(:accounts_users, [:email], unique: true)
  end
end
