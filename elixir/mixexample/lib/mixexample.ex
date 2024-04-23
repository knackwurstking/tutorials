# defmodule Mixexample do
#  @moduledoc """
#  Documentation for `Mixexample`.
#
#  Compile with `mix compile`
#  Run with `mix -e "Mixexample.hello"`
#  Or with iex `iex -S mix` then run `Mixexample.hello`
#  """
#
#  @doc """
#  Hello world.
#
#  ## Examples
#
#      iex> Mixexample.hello()
#      :world
#
#  """
#  def hello do
#    IO.puts(:world)
#  end
# end

defmodule Mixexample do
  use Application
  alias UUID

  def start(_type, _args) do
    IO.puts(UUID.uuid4())
    Supervisor.start_link([], strategy: :one_for_one)
  end

  def hello do
    :world
  end
end
