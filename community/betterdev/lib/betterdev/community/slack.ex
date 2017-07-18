defmodule Betterdev.Community.SlackRtm do
  use Slack
  alias Betterdev.Community.Bot

  def handle_connect(slack, state) do
    IO.puts "Connected as #{slack.me.name}"
    {:ok, state}
  end

  def handle_event(message = %{type: "message"}, slack, state) do
    IO.inspect message
    Bot.import_link(message.text)
    send_message("Thanks. I have imported it. View it on https://one.betterdev.link", message.channel, slack)
    {:ok, state}
  end
  def handle_event(_, _, state), do: {:ok, state}

  def handle_info({:message, text, channel}, slack, state) do
    IO.puts "Sending your message, captain!"

    send_message(text, channel, slack)

    {:ok, state}
  end
  def handle_info(_, _, state), do: {:ok, state}
end
