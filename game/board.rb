require_relative 'game'
require_relative 'snake'
require_relative 'food'

class Board
  attr_reader :game
  attr_reader :data

  def initialize(game, data)
    @game = game
    @data = data
  end

  %w[height width food snakes].each do |key|
    define_method("#{key}_data") { data[key] }
  end

  alias width width_data
  alias height height_data

  def [](x, y)
    return nil unless valid_xy(x, y)
    grid[y][x]
  end

  def []=(x, y, value)
    raise 'board assignment out of bounds' unless valid_xy(x, y)
    grid[y][x] = value
  end

  def valid_xy(x, y)
    (0...width).cover?(x) && (0...height).cover?(y)
  end

  def food
    @food ||= food_data.map { |food| Food.new(self, food['x'], food['y']) }
  end

  def snakes
    @snakes ||= snakes_data.map { |snake| Snake.new(self, snake) }.sort_by(&:id)
  end

  def rows
    grid
  end

  def columns
    grid.transpose
  end

  def to_s
    builder = StringIO.new

    rows.each do |row|
      row.each do |position|
        builder << position
        builder << ' '
      end

      builder << "\n"
    end

    builder.string
  end

  private

  def grid
    @grid ||= Array.new(width) { |y| Array.new(height) { |x| Position.new(self, x, y) } }
  end
end

class Position
  attr_reader :board
  attr_reader :x
  attr_reader :y
  attr_accessor :segments
  attr_accessor :food

  alias food? food

  def initialize(board, x, y, segments = [], food = nil)
    @board = board
    @x = x
    @y = y
    @segments = segments
    @food = food
  end

  def snakes
    @snakes ||= segments.map(&:snake)
  end

  def distance(other)
    [x - other.x, y - other.y].map(&:abs).sum
  end

  def up(d = 1)
    [x, y - d]
  end

  def down(d = 1)
    [x, y + d]
  end

  def left(d = 1)
    [x - d, y]
  end

  def right(d = 1)
    [x + d, y]
  end

  def to_s
    ([food ? 'F' : '-'] + board.snakes.map { |s| snakes.include?(s) ? (s.you? ? 'Y' : 'S') : '-' }).join
  end
end
