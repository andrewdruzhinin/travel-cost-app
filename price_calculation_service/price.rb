class Price
  attr_reader :distance
  attr_reader :duration
  attr_reader :config

  def initialize(distance, duration)
    @distance = distance
    @duration = duration
    @config = YAML.load_file('config.yml')
  end

  def calc
    distance_price = (distance / 1000.0) * config['per_km']
    duration_price = (duration / 60.0) * config['per_minute']

    price = (config['supply'] + duration_price + distance_price).round

    price < config['min_price'] ? config['min_price'] : price
  end
end
