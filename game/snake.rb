require_relative 'game'
require_relative 'board'

class Snake
  attr_reader :board
  attr_reader :data

  def initialize(board, data)
    @board = board
    @data = data
  end

  %w[id name health body].each do |key|
    define_method("#{key}_data") { data[key] }
  end

  alias id id_data
  alias name name_data
  alias health health_data

  def you?
    @you ||= self == board.game.you
  end

  def ==(other)
    id == other.id
  end

  def body
    @body ||= body_data.map { |body| Segment.new(self, body['x'], body['y']) }
  end
  alias segments body

  def head
    body.first
  end

  def enemy?
    !you?
  end

  def killable?
    enemy? && board.game.you.size > segments.size
  end

  def index
    @index ||= board.snakes.index { |snake| snake.id == id }
  end

  def inspect
    "<#{you? ? 'You' : 'Snake'} id=#{id} name=#{name} health=#{health} x=#{head.position.x} y=#{head.position.y}>"
  end
end

class Segment
  attr_reader :snake
  attr_reader :position

  def initialize(snake, x, y)
    @snake = snake
    @position = snake.board[x, y]

    position.segments << self
  end

  def head?
    self == snake.head
  end

  def inspect
    "<Segment snake_id=#{snake.id} x=#{position.x} y=#{position.y}>"
  end
end
