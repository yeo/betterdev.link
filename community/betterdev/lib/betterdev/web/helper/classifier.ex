defmodule Betterdev.Helper.Classifier do
  @taxonomy [
    {:ruby, ["ruby", ".rb", "rails", "active record", "tdd", "def", "module", "class", "gem"]},
    {:php, ["php", ".php", "laravel", "symfony", "composer", "packagist"]},
    {:go, ["go lang", ".go", "golang", "type", "struct"]},
    {:elixir, ["elixir", ".ex", "iex", "mix", "hex"]},
    {:erlang, ["erlang", ".erl", "rebar", "erl", "IO", "mix"]},
    {:nodejs, ["nodejs", "javascript", ".js", "v8", "es6", "jsx"]},
    {:linux, ["linux", "bash", "vim", "terminal", "server", "system", "configuration", "gdb", "ssh"]},
    {:algorithm, ["algorithm", "code", "interview", "programming", "tree", "stack", "queue", "question", "job", "crawl", "scrape", "structure"]},
    {:design, ["design", "css", "typography", "typeface"]},
  ]

  def extract(content) do
    Enum.flat_map(@taxonomy, fn ({term, words}) ->
      IO.inspect term
      IO.inspect words
      Enum.reduce(words, 0, fn (w, acc) ->
        IO.inspect w
        IO.inspect acc
        case String.contains?(content, w) do
          true -> acc + 1
          false -> acc
        end
      end)
    end)
  end
end
