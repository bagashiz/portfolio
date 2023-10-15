/**
 * Experience data for the resume.
 *
 * @type {Array<{
 *   job: string,
 *   url: string,
 *   company: string,
 *   start: string,
 *   end: string,
 *   description: Array<string>
 * }>}
 */
export const experiences = [
	{
		job: 'Head of Public Relations Division',
		url: 'https://porosfilkom.ub.ac.id',
		company: 'POROS Organization of Open Source',
		start: 'February 2023',
		end: 'Present',
		description: [
			'Established and maintained relationships with external groups of interest, resulting in higher participation and support for organization activities.',
			"Managed organization's social media presence, resulting in increased engagement and promotion of events and initiatives.",
			'Utilized digital platforms to enhance brand recognition and advance the objectives of the organization.'
		]
	},
	{
		job: 'Laboratory Teaching Assistant of Database Administration',
		url: 'https://filkom.ub.ac.id',
		company: 'Faculty of Computer Science (FILKOM) Brawijaya University',
		start: 'August 2023',
		end: 'Present',
		description: [
			"Led lab activities of database administration using IBM DB2 that applied lecturer's theory and developed students' practical skills.",
			'Provided support and guidance to help students understand Linux system administration and database administration concepts using IBM DB2.',
			"Designed and graded assignments and projects to assess students' understanding of database administration concepts using IBM DB2."
		]
	},
	{
		job: 'Laboratory Teaching Assistant of Advanced Programming',
		url: 'https://filkom.ub.ac.id',
		company: 'Faculty of Computer Science (FILKOM) Brawijaya University',
		start: 'February 2023',
		end: 'June 2023',
		description: [
			"Led lab activities of OOP using Java that applied lecturer's theory and developed students' practical skills.",
			'Provided support and guidance to help students understand OOP concepts using Java.',
			"Designed and graded programming assignments and projects to assess students' understanding of OOP concepts using Java."
		]
	},
	{
		job: 'Project-Based Intern: Backend Developer Virtual Internship Experience Evermos',
		url: 'https://www.rakamin.com/virtual-internship-experience/back-end-developer-evermos',
		company: 'Rakamin Academy',
		start: 'April 2023',
		end: 'May 2023',
		description: [
			'Created API services using the Go programming language.',
			'Built and managed integration between API services and MySQL database.',
			'Tested API functionality using Postman API to ensure service reliability.'
		]
	},
	{
		job: 'Staff of Public Relations Division',
		url: 'https://porosfilkom.ub.ac.id',
		company: 'POROS Organization of Open Source',
		start: 'February 2022',
		end: 'January 2023',
		description: [
			'Established and maintained relationships with external groups of interest, resulting in higher participation and support for organization activities.',
			"Managed organization's social media presence, resulting in increased engagement and promotion of events and initiatives.",
			'Utilized digital platforms to enhance brand recognition and advance the objectives of the organization.'
		]
	}
];

/**
 * Volunteering data for the resume.
 *
 * @type {Array<{
 *   job: string,
 *   url: string,
 *   company: string,
 *   start: string,
 *   end: string,
 *   description: Array<string>
 * }>}
 */
export const volunteering = [
	{
		job: 'Staff of IT Division',
		url: 'https://hology.ub.ac.id',
		company: 'Hology UB 6.0',
		start: 'May 2023',
		end: 'Present',
		description: [
			'Collaborated with back-end team to develop robust APIs using TypeScript and Express.js for the event website.',
			'Managed CI/CD pipeline using GitHub Actions and GitHub Container Registry to efficiently package and release back-end APIs.',
			'Utilized Docker containers and Nginx to deploy back-end APIs, optimizing scalability and enhancing reliability.',
			'Managed back-end and front-end server maintenance to ensure uninterrupted performance during the event.'
		]
	}
];

/**
 * Social data for the resume.
 *
 * @type {Array<{
 *   href: string,
 *   rel: string,
 *   fa_class: string,
 * }>}
 */
export const socials = [
	{
		href: 'https://linkedin.com/in/bagas-hizbullah',
		rel: '',
		fa_class: 'linkedin-in'
	},
	{
		href: 'https://github.com/bagashiz',
		rel: '',
		fa_class: 'github'
	},
	{
		href: 'https://dev.to/bagashiz',
		rel: '',
		fa_class: 'dev'
	},
	{
		href: 'https://mastodon.social/@Pak_Dengklek',
		rel: 'me',
		fa_class: 'mastodon'
	}
];

/**
 * Education data for the resume.
 *
 * @type {Array<{
 *  school: string,
 * url: string,
 * major: string,
 * start: string,
 * end: string,
 * description: string
 * }>}
 */
export const educations = [
	{
		school: 'Brawijaya University',
		url: 'https://ub.ac.id',
		major: 'Bachelor of Information Systems',
		start: 'August 2021',
		end: 'Present',
		description: 'GPA: 3.89/4.00'
	},
	{
		school: 'SMAN 1 Bekasi',
		url: 'https://sman1bekasi.sch.id',
		major: 'Mathematics and Natural Sciences',
		start: 'July 2018',
		end: 'May 2021',
		description: 'GPA: 90.73/100'
	}
];

/** Font Awesome icons name for the skills data. */
export const skills = [
	'golang',
	'java',
	'html5',
	'css3-alt',
	'js-square',
	'node-js',
	'php',
	'laravel',
	'npm',
	'git-alt',
	'linux',
	'docker',
	'digital-ocean',
	'aws'
];

/** Workflows data for the resume. */
export const workflows = ['Clean Architecture', 'Test-Driven Development', 'Data-Driven Testing'];

/**
 * Awards data for the resume.
 *
 * @type {Array<{
 * place: number,
 * suffix: string,
 * host: string,
 * competition: string,
 * translation: string
 * }>}
 */
export const awards = [
	{
		place: 1,
		suffix: 'st',
		host: 'Brawijaya University',
		competition: 'Lomba Penulisan Proposal Program Kreativitas Mahasiswa (LP2PKM) 2022',
		translation: 'Student Creativity Program Proposal Writing Competition 2022'
	},
	{
		place: 1,
		suffix: 'st',
		host: 'Medan State University',
		competition: 'Business Plan Competition Nasional UNIMED 2021',
		translation: 'UNIMED National Business Plan Competition 2021'
	}
];

/**
 * Certification data for the resume.
 *
 * @type {Array<{
 *  title: string,
 * credential_id: string,
 * credential_url: string
 * }>}
 */
export const certifications = [
	{
		title: 'AWS Cloud Technical Essentials',
		credential_id: 'HSKQ9QM64B5A',
		credential_url: 'https://www.coursera.org/account/accomplishments/certificate/HSKQ9QM64B5A'
	},
	{
		title: 'Belajar Fundamental Aplikasi Back-End',
		credential_id: '98XWV6M3WPM3',
		credential_url: 'https://www.dicoding.com/certificates/98XWV6M3WPM3'
	},
	{
		title: 'Architecting on AWS (Membangun Arsitektur Cloud di AWS)',
		credential_id: 'EYX46KOVWPDL',
		credential_url: 'https://www.dicoding.com/certificates/EYX46KOVWPDL'
	},
	{
		title: 'Belajar Membuat Aplikasi Back-End untuk Pemula',
		credential_id: 'JLX1L9V96X72',
		credential_url: 'https://www.dicoding.com/certificates/JLX1LVQ32X72'
	},
	{
		title: 'Belajar Dasar Pemrograman JavaScript',
		credential_id: 'JLX1L9V96X72',
		credential_url: 'https://www.dicoding.com/certificates/JLX1LVQ32X72'
	},
	{
		title: 'Cloud Practitioner Essentials (Belajar Dasar AWS Cloud)',
		credential_id: 'JLX1L9V96X72',
		credential_url: 'https://www.dicoding.com/certificates/JLX1LVQ32X72'
	},
	{
		title: 'Belajar Dasar Structured Query Language (SQL)',
		credential_id: '2VX3J6603PYQ',
		credential_url: 'https://www.dicoding.com/certificates/2VX3J6603PYQ'
	},
	{
		title: 'Backend Master Class [Golang + Postgres + Kubernetes + gRPC]',
		credential_id: 'UC-69e6c3c2-7a66-4744-8b60-c47ea0702897',
		credential_url: 'https://www.udemy.com/certificate/UC-69e6c3c2-7a66-4744-8b60-c47ea0702897'
	},
	{
		title: 'Introduction to IoT',
		credential_id: 'cdf42e7d-a405-47b6-aa5c-b8b66f360f37',
		credential_url: 'https://www.credly.com/badges/cdf42e7d-a405-47b6-aa5c-b8b66f360f37'
	}
];
