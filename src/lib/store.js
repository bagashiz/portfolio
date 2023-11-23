/**
 * Work experience data for the resume.
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
export const works = [
	{
		job: 'Laboratory Teaching Assistant of Database Administration',
		url: 'https://filkom.ub.ac.id',
		company: 'Faculty of Computer Science (FILKOM) Brawijaya University',
		start: 'August 2023',
		end: 'Present',
		description: [
			'Enhanced the practical skills and comprehension of database administration principles using IBM DB2 for over 30 students by leading more than 10 hands-on lab sessions.',
			'Strengthened the understanding of Linux systems and key concepts of database administration for students through guidance and support.',
			"Evaluate students' proficiency in database administration using IBM DB2 by designing well-structured assessments."
		]
	},
	{
		job: 'Project-Based Intern: Back-End Developer Virtual Internship Experience Evermos',
		url: 'https://www.rakamin.com/virtual-internship-experience/back-end-developer-evermos',
		company: 'Rakamin Academy',
		start: 'April 2023',
		end: 'May 2023',
		description: [
			'Created robust web API services using the Go programming language by implementing clean architecture principles.',
			'Optimized the overall functionality of the system through seamless integration between API services and MySQL database.',
			'Ensured stability and reliability of the services through rigorous API testing using Postman API.'
		]
	},
	{
		job: 'Laboratory Teaching Assistant of Advanced Programming',
		url: 'https://filkom.ub.ac.id',
		company: 'Faculty of Computer Science (FILKOM) Brawijaya University',
		start: 'February 2023',
		end: 'June 2023',
		description: [
			'Enhanced the practical skills and comprehension of object-oriented programming using Java for over 20 students by leading more than 10 hands-on lab sessions.',
			'Strengthened the understanding of Java libraries and key concepts of object-oriented programming for students through guidance and support.',
			"Evaluate students' proficiency in object-oriented programming using Java by designing well-structured assessments."
		]
	}
];

/**
 * Organization experience data for the resume.
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
export const organizations = [
	{
		job: 'Web & Cloud Curriculum Mentor',
		url: 'https://gdsc.community.dev/university-of-brawijaya',
		company: 'Google Developer Student Clubs University of Brawijaya Chapter',
		start: 'October 2023',
		end: 'Present',
		description: [
			'Developed and shared learning materials on web and cloud development to more than 100 members by collaborating with fellow mentors.',
			'Provided support and guidance to help members learn web and cloud development through mentorship and feedback.',
			'Delivered insights on web and cloud development as a workshop speaker to all members.'
		]
	},
	{
		job: 'Chief of Public Relations Division',
		url: 'https://porosfilkom.ub.ac.id',
		company: 'POROS Organization of Open Source',
		start: 'February 2023',
		end: 'Present',
		description: [
			'Increased participation and support for organization activities by establishing and strengthening relationships with external groups of interest.',
			'Improved brand recognition and progressed organizational objectives by leveraging digital platforms.',
			"Increased event engagement and promotion up to 15% by monitoring the organization's social media presence."
		]
	},
	{
		job: 'DevOps Staff of IT Division',
		url: 'https://hology.ub.ac.id',
		company: 'Hology UB 6.0',
		start: 'May 2023',
		end: 'October 2023',
		description: [
			'Packaged and released 2 back-end web API services efficiently by managing CI/CD pipeline using GitHub Actions and GitHub Container Registry.',
			'Optimized scalability and enhanced reliability of the web API services deployment by utilizing Docker containers and Nginx web proxy.',
			'Ensured uninterrupted performance of the web by managing back-end and front-end server maintenance.'
		]
	},
	{
		job: 'Staff of Public Relations Division',
		url: 'https://porosfilkom.ub.ac.id',
		company: 'POROS Organization of Open Source',
		start: 'February 2022',
		end: 'January 2023',
		description: [
			'Increased participation and support for organization activities by establishing and strengthening relationships with external groups of interest.',
			'Improved brand recognition and progressed organizational objectives by leveraging digital platforms.',
			"Increased event engagement and promotion up to 10% by monitoring the organization's social media presence."
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
		description: 'Cumulative GPA of 3.89/4.00'
	},
	{
		school: 'SMAN 1 Bekasi',
		url: 'https://sman1bekasi.sch.id',
		major: 'Mathematics and Natural Sciences',
		start: 'July 2018',
		end: 'May 2021',
		description: 'Cumulative GPA of 90.73/100'
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

/** Interests paragraphs for the resume. */
export const interests = [
	'Besides being a busy information systems student, I enjoy using my free time to stay informed about the latest developments of Free and Open Source Software (FOSS) projects, with a particular interest in the Linux Desktop ecosystem. I participate in online forums to discuss these topics with others and share my thoughts. I also enjoy writing blog posts on subjects that interest me. Additionally, I constantly seek out opportunities to learn and improve my skills through courses and workshops I stumble upon online.',
	'I am currently learning back-end web development, cloud computing, and software architecture. These skills are important for creating effective, efficient, and secure systems. Having these knowledge will also allow me to understand how to create and implement more complex systems.'
];

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
		title: 'Architecting on AWS (Membangun Arsitektur Cloud di AWS)',
		credential_id: 'EYX46KOVWPDL',
		credential_url: 'https://www.dicoding.com/certificates/EYX46KOVWPDL'
	},
	{
		title: 'Cloud Practitioner Essentials (Belajar Dasar AWS Cloud)',
		credential_id: 'JLX1L9V96X72',
		credential_url: 'https://www.dicoding.com/certificates/JLX1LVQ32X72'
	},
	{
		title: 'Belajar Fundamental Aplikasi Back-End',
		credential_id: '98XWV6M3WPM3',
		credential_url: 'https://www.dicoding.com/certificates/98XWV6M3WPM3'
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
		title: 'Belajar Dasar Structured Query Language (SQL)',
		credential_id: '2VX3J6603PYQ',
		credential_url: 'https://www.dicoding.com/certificates/2VX3J6603PYQ'
	},
	{
		title: 'Back-End Master Class [Golang + Postgres + Kubernetes + gRPC]',
		credential_id: 'UC-69e6c3c2-7a66-4744-8b60-c47ea0702897',
		credential_url: 'https://www.udemy.com/certificate/UC-69e6c3c2-7a66-4744-8b60-c47ea0702897'
	},
	{
		title: 'Introduction to IoT',
		credential_id: 'cdf42e7d-a405-47b6-aa5c-b8b66f360f37',
		credential_url: 'https://www.credly.com/badges/cdf42e7d-a405-47b6-aa5c-b8b66f360f37'
	}
];
