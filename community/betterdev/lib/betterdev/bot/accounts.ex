defmodule Betterdev.Bot.Accounts do
  @moduledoc """
  The boundary for the Accounts system.
  """

  import Ecto.Query, warn: false
  alias Betterdev.Repo

  alias Betterdev.Accounts.User

  def retreive_from_bot_user(channel, uuid, ) do
    case Repo.get_by(Betterdev.Accounts.User, channel_user_id: uuid, channel: channel) do
      nil ->
        {:ok, u} = Betterdev.Accounts.create_user(%{channel: channel, channel_user_id: uuid})
        u
      u -> u
    end
  end
end
