defmodule Betterdev.Community.Bot do
  alias Betterdev.Account
  alias Betterdev.Community.Link

  alias Exbot.Request.Update
  alias Betterdev.Repo


  def start do
    listen(-100)
    #Enum.map(links , &(Link.changeset(%Link{}, %{user: import_user, title: &1, uri: &1}) |> Repo.insert()))
    #Enum.map(links , &(Link.changeset(%Link{}, %{user: import_user, title: &1, uri: &1}) |> Repo.insert()))
  end

  def import_user do
    user = Accounts.get_user!(1)
  end

  def import_link(text) do
    IO.puts "Import #{text}"
    String.split(text, "\n") |> Enum.map(fn (line) ->
      case Regex.run(~r/(http|https):\/\/([^\s\t\n]+)/, text, global: true) do
        [url | _] ->
          w = Scrape.website(url)
          if w.title do
            %Link{user_id: 1, title: w.title || url, uri: url, description: w.description, picture: w.image || w.favicon, status: "published", } |> Repo.insert()
          end
				_ -> IO.puts "#{text} has no url"
      end
    end)
  end

  def listen(start \\ -100) do
    IO.puts "Listen with #{start}"
    messages =  Exbot.get_updates(&(&1 |> Update.with_offset(start)))
    last_message = messages |> List.last
    messages |> Enum.filter_map(&(&1.message && &1.message.chat.id == -1001120191455 && &1.message.text), &(import_link(&1.message.text)))
    next = if last_message == nil do
      :timer.sleep(10000)
      if start > 0 do
        next = start
      else
        next = -100
      end
    else
      last_message.update_id + 1
    end
    listen(next)
  end
end
