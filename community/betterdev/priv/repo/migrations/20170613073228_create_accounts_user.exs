defmodule Betterdev.Repo.Migrations.CreateBetterdev.Accounts.User do
  use Ecto.Migration

  def change do
    create table(:accounts_users) do
      add :email, :string
      add :name, :string

      timestamps()
    end

  end
end
