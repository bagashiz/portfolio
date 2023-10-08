/**
 * Represents a blog object.
 *
 * @typedef {Object} Blog
 * @property {string} type_of - The type of the blog.
 * @property {number} id - The unique identifier of the blog.
 * @property {string} title - The title of the blog.
 * @property {string} description - The description of the blog.
 * @property {string} readable_publish_date - The readable publish date of the blog.
 * @property {string} slug - The slug of the blog.
 * @property {string} path - The path of the blog.
 * @property {string} url - The URL of the blog.
 * @property {number} comments_count - The number of comments on the blog.
 * @property {number} public_reactions_count - The number of public reactions on the blog.
 * @property {number} collection_id - The collection ID of the blog.
 * @property {string} published_timestamp - The timestamp when the blog was published.
 * @property {number} positive_reactions_count - The number of positive reactions on the blog.
 * @property {string} cover_image - The cover image URL of the blog.
 * @property {string} social_image - The social image URL of the blog.
 * @property {string} canonical_url - The canonical URL of the blog.
 * @property {string} created_at - The timestamp when the blog was created.
 * @property {string} edited_at - The timestamp when the blog was edited.
 * @property {string} crossposted_at - The timestamp when the blog was crossposted.
 * @property {string} published_at - The timestamp when the blog was published.
 * @property {string} last_comment_at - The timestamp of the last comment on the blog.
 * @property {number} reading_time_minutes - The reading time of the blog in minutes.
 * @property {string[]} tag_list - The list of tags associated with the blog.
 * @property {string} tags - The tags of the blog as a comma-separated string.
 * @property {BlogUser} user - The user who created the blog.
 */
type Blog = {
	type_of: string;
	id: number;
	title: string;
	description: string;
	readable_publish_date: string;
	slug: string;
	path: string;
	url: string;
	comments_count: number;
	public_reactions_count: number;
	collection_id: number;
	published_timestamp: string;
	positive_reactions_count: number;
	cover_image: string;
	social_image: string;
	canonical_url: string;
	created_at: string;
	edited_at: string;
	crossposted_at: string;
	published_at: string;
	last_comment_at: string;
	reading_time_minutes: number;
	tag_list: string[];
	tags: string;
	user: BlogUser;
};

/**
 * Represents a user associated with a blog.
 *
 * @typedef {Object} BlogUser
 * @property {string} name - The name of the user.
 * @property {string} username - The username of the user.
 * @property {string} twitter_username - The Twitter username of the user.
 * @property {string} github_username - The GitHub username of the user.
 * @property {string} website_url - The website URL of the user.
 * @property {string} profile_image - The profile image URL of the user.
 * @property {string} profile_image_90 - The profile image URL (90x90) of the user.
 */
type BlogUser = {
	name: string;
	username: string;
	twitter_username: string;
	github_username: string;
	website_url: string;
	profile_image: string;
	profile_image_90: string;
};

/**
 * Represents a GitHub repository with its details.
 *
 * @typedef {Object} Project
 * @property {string} name - The name of the project.
 * @property {string} description - The description of the project.
 * @property {string} url - The URL of the project.
 * @property {Stargazer} stargazers - The stargazers details of the project.
 * @property {Fork} forks - The forks details of the project.
 */
type Project = {
	name: string;
	description: string;
	url: string;
	stargazers: Stargazer;
	forks: Fork;
};

/**
 * Represents the stargazers details of a project.
 *
 * @typedef {Object} Stargazer
 * @property {number} totalCount - The total count of stargazers.
 */
type Stargazer = {
	totalCount: number;
};

/**
 * Represents the forks details of a project.
 *
 * @typedef {Object} Fork
 * @property {number} totalCount - The total count of forks.
 */
type Fork = {
	totalCount: number;
};
