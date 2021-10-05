if redis.call('exists', KEYS[1]) == 0
then
	redis.call('set', KEYS[1], ARGV[2]);
	redis.call('publish', KEYS[2], ARGV[1]);
	return 1
else
	return 0
end