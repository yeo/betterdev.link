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
    {:css, ["css", "html"]},
  ]

  @super_taxonomy [
    {:ruby, ["ruby", ".rb", "rails", "gem", "hanami", "basecamp", "dhh", "tenderlove", "jruby", "rvm", "rbenv", "sinatra"]},
    {:php, ["php", ".php", "laravel", "symfony", "composer", "packagist", "hhvm"]},
    {:go, ["go lang", ".go", "golang", "type", "struct", "interface", "fmt", "net/http"]},
    {:elixir, ["elixir", ".ex", "iex", "mix", "hex", "ecto", "phoenix", "plug"]},
    {:erlang, ["erlang", ".erl", "rebar", "erl", "IO", "mix", "cowboy",]},
    {:nodejs, ["nodejs", "javascript", ".js", "v8", "es6", "jsx", "react", "mithril", "express", "promise"]},
    {:linux, ["linux", "bash", "vim", "terminal", "server", "system", "configuration", "gdb", "ssh", "systemd", ]},
    {:algorithm, ["algorithm", "code", "interview", "programming", "tree", "stack", "queue", "question", "job", "crawler", "scrape", "structure", "tdd"]},
    {:design, ["design", "css", "typography", "typeface", "font", "bootstrap", "foundation", "purecss", "svg", "icon", "fontawesome", "webfont", "ttf", "otf", "flexbox", "float"]},
    {:css, ["css", "html", "css3", "flexbox", "bootstrap", "purecss"]},
    {:devops, ["aws", "cloudformation", "ansible", "chef", "puppet", "cloudfront", "cloudflare", "nginx", "haproxy", "cache"]},
    {:database, ["mysql", "postgres", "sql", "mongodb", "rethinkdb", "query", "index", "group by", "table", ]},
    {:python, ["py", "python", "django", "flask", "jinja", "python3", "pip", "easy_install", "pypi", "bottle"]},
  ]

  def extract(terms, threshold, content) do
    Enum.map(terms, fn ({term, words}) ->
      score = Enum.reduce(words, 0, fn (w, acc) ->
        case String.contains?(content, w) do
          true -> acc + 1
          false -> acc
        end
      end)
      {term, score}
    end) |> Enum.filter_map(fn ({term, score}) -> score >= threshold end, fn({term, score}) -> term end)
  end

  def extract(content) do
    tags = extract(@super_taxonomy, 2, content)
    unsure_tags = extract(@taxonomy, 3, content)
    tags ++ unsure_tags
  end
end
