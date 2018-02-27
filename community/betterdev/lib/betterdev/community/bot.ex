defmodule Betterdev.Community.Bot do
  use GenServer

  @interval 30000

  alias Betterdev.Account
  alias Betterdev.Community.Link

  alias Exbot.Request.Update
  alias Betterdev.Repo
  alias Betterdev.Bot.Accounts, as: BotUser

  def start_link do
    :gen_server.start_link(__MODULE__, [], [])
  end

  def init(state) do
    schedule_work()
    {:ok, -100}
  end

  def schedule_work do
    Process.send_after(self(), :work, @interval)
  end

  def handle_info(:work, state) do
    next = listen(state)
    schedule_work() # Reschedule once more
    {:noreply, next}
  end

  def import_user, do: user = Accounts.get_user!(1)
  def import_user(u), do: BotUser.retreive_from_bot_user(u)

  def import_link(text, user) do
    String.split(text, "\n") |> Enum.map(fn (line) ->
      case Regex.run(~r/(http|https):\/\/([^\s\t\n]+)/, text, global: true) do
        [url | _] ->
          w = Scrape.website(url)
          if w.title do
            link = %Link{user: import_user(user), title: w.title || url, uri: url, description: w.description, picture: w.image || w.favicon, status: "published", } |> Repo.insert()
            Task.start_link(fn -> Community.post_process_link(link) end)
          end
          _ -> IO.puts "#{text} has no url"
      end
    end)
  end

  # def retri
  def listen(start \\ -100) do
    messages =  Exbot.get_updates(&(&1 |> Update.with_offset(start)))
    last_message = messages |> List.last

    try do
      messages |> Enum.filter_map(&(&1.message && &1.message.text), &(import_link(&1.message.text, &1.message.from)))
    rescue
      # TODO: We may want to report to some monitoring system
      # We don't want it to break this process. Hence ignore error and continue
      e in RuntimeError -> IO.inspect e
    end

    if last_message == nil do
      if start > 0 do
        next = start
      else
        next = -100
      end
    else
      last_message.update_id + 1
    end
  end
end
