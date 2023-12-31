import { env } from '$env/dynamic/private';
import { Redis } from 'ioredis';

/**
 * @type {Redis} redis - Redis client instance
 */
const redis = new Redis({
	host: env.REDIS_HOST,
	port: parseInt(env.REDIS_PORT || '6379')
});

redis.on('error', (error) => {
	console.error(`Error initializing redis client: ${error}`);
});

export default redis;
