defmodule Betterdev.Repo.Migrations.AddBotChannelToUser do
  use Ecto.Migration

  def change do
    alter table(:accounts_users) do
      add :channel, :string
      add :channel_user_id, :string
    end
  end
end
