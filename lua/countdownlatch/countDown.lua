local v = redis.call('decr', KEYS[1])

if v <= 0
then
    redis.call('del', KEYS[1])
end

if v == 0
then
    redis.call('publish', KEYS[2], ARGV[1])
end;