package model

var (
	// Socials is a slice of Social struct that contains list of social media profiles.
	Socials = []Social{
		{
			Name:    "LinkedIn",
			Href:    "https://linkedin.com/in/bagas-hizbullah",
			Rel:     "",
			FaClass: "linkedin-in",
		},
		{
			Name:    "GitHub",
			Href:    "https://github.com/bagashiz",
			Rel:     "",
			FaClass: "github",
		},
		{
			Name:    "Dev",
			Href:    "https://dev.to/bagashiz",
			Rel:     "",
			FaClass: "dev",
		},
		{
			Name:    "Mastodon",
			Href:    "https://hachyderm.io/@bagashiz",
			Rel:     "me",
			FaClass: "mastodon",
		},
	}

	// Works is a slice of Experience struct that contains list of work experiences.
	Works = []Experience{
		{
			Job:     "Junior Back-End Engineer Intern",
			Url:     "https://godentist.co.id/",
			Company: "GoDentist (PT PERIKSA GIGI INDONESIA)",
			Start:   "February 2024",
			End:     "Present",
			Descriptions: []string{
				"Developed, maintained, and co-designed web service APIs using JavaScript and Node.js.",
				"Troubleshooted and resolved API issues to maintain optimal performance and user experience.",
				"Collaborated effectively with cross-functional stakeholders to ensure alignment with business requirements.",
			},
		},
		{
			Job:     "Laboratory Teaching Assistant of Database Administration",
			Url:     "https://filkom.ub.ac.id",
			Company: "Faculty of Computer Science (FILKOM) Brawijaya University",
			Start:   "August 2023",
			End:     "December 2023",
			Descriptions: []string{
				"Enhanced the practical skills and comprehension of database administration principles using IBM DB2 for over 30 students by leading more than 10 hands-on lab sessions.",
				"Strengthened the understanding of Linux systems and key concepts of database administration for students through guidance and support.",
				"Evaluate students' proficiency in database administration using IBM DB2 by designing well-structured assessments.",
			},
		},
		{
			Job:     "Laboratory Teaching Assistant of Advanced Programming",
			Url:     "https://filkom.ub.ac.id",
			Company: "Faculty of Computer Science (FILKOM) Brawijaya University",
			Start:   "February 2023",
			End:     "June 2023",
			Descriptions: []string{
				"Enhanced the practical skills and comprehension of object-oriented programming using Java for over 20 students by leading more than 10 hands-on lab sessions.",
				"Strengthened the understanding of Java libraries and key concepts of object-oriented programming for students through guidance and support.",
				"Evaluate students' proficiency in object-oriented programming using Java by designing well-structured assessments.",
			},
		},
		{
			Job:     "Project-Based Intern: Back-End Developer Virtual Internship Experience Evermos",
			Url:     "https://www.rakamin.com/virtual-internship-experience/back-end-developer-evermos",
			Company: "Rakamin Academy",
			Start:   "April 2023",
			End:     "May 2023",
			Descriptions: []string{
				"Created robust web API services using the Go programming language by implementing clean architecture principles.",
				"Optimized the overall functionality of the system through seamless integration between API services and MySQL database.",
				"Ensured stability and reliability of the services through rigorous API testing using Postman API.",
			},
		},
	}

	// Volunteers is a slice of Experience struct that contains list of volunteer experiences.
	Volunteers = []Experience{
		{
			Job:     "Vice Chairman",
			Url:     "https://porosfilkom.ub.ac.id",
			Company: "POROS Organization of Open Source",
			Start:   "February 2024",
			End:     "Present",
			Descriptions: []string{
				"Acted as a lead substitute for the Chairman, managing essential organizational activities.",
				"Bridged internal and external communication gaps by acting as a liaison between leadership, members, and other groups of interest.",
				"Provided strategic guidance and support to other members, fostering a collaborative and productive work environment.",
			},
		},
		{
			Job:     "Web & Cloud Curriculum Mentor",
			Url:     "https://gdsc.community.dev/university-of-brawijaya",
			Company: "Google Developer Student Clubs University of Brawijaya Chapter",
			Start:   "October 2023",
			End:     "Present",
			Descriptions: []string{
				"Developed and shared learning materials on web and cloud development to more than 100 members by collaborating with fellow mentors.",
				"Provided support and guidance to help members learn web and cloud development through mentorship and feedback.",
				"Delivered insights on web and cloud development as a workshop speaker to all members.",
			},
		},
		{
			Job:     "Chief of Public Relations Division",
			Url:     "https://porosfilkom.ub.ac.id",
			Company: "POROS Organization of Open Source",
			Start:   "February 2023",
			End:     "January 2024",
			Descriptions: []string{
				"Increased participation and support for organization activities by establishing and strengthening relationships with external groups of interest.",
				"Improved brand recognition and progressed organizational objectives by leveraging digital platforms.",
				"Increased event engagement and promotion up to 15% by monitoring the organization's social media presence.",
			},
		},
		{
			Job:     "DevOps Staff of IT Division",
			Url:     "https://hology.ub.ac.id",
			Company: "Hology UB 6.0",
			Start:   "May 2023",
			End:     "October 2023",
			Descriptions: []string{
				"Packaged and released 2 back-end web API services efficiently by managing CI/CD pipeline using GitHub Actions and GitHub Container Registry.",
				"Optimized scalability and enhanced reliability of the web API services deployment by utilizing Docker containers and Nginx web proxy.",
				"Ensured uninterrupted performance of the web by managing back-end and front-end server maintenance.",
			},
		},
		{
			Job:     "Staff of Public Relations Division",
			Url:     "https://porosfilkom.ub.ac.id",
			Company: "POROS Organization of Open Source",
			Start:   "February 2022",
			End:     "January 2023",
			Descriptions: []string{
				"Increased participation and support for organization activities by establishing and strengthening relationships with external groups of interest.",
				"Improved brand recognition and progressed organizational objectives by leveraging digital platforms.",
				"Increased event engagement and promotion up to 10% by monitoring the organization's social media presence.",
			},
		},
	}

	// Educations is a slice of Education struct that contains list of educational backgrounds.
	Educations = []Education{
		{
			School:      "Brawijaya University",
			Url:         "https://ub.ac.id",
			Major:       "Bachelor of Information Systems",
			Start:       "August 2021",
			End:         "Present",
			Description: "Cumulative GPA of 3.90/4.00",
		},
		{
			School:      "Bangkit Academy led by Google, Tokopedia, Gojek, & Traveloka",
			Url:         "https://bangkit.academy",
			Major:       "Cloud Computing",
			Start:       "February 2024",
			End:         "Present",
			Description: "Indepedent Study Program of Kampus Merdeka",
		},
		{
			School:      "SMAN 1 Bekasi",
			Url:         "https://sman1bekasi.sch.id",
			Major:       "Mathematics and Natural Sciences",
			Start:       "July 2018",
			End:         "May 2021",
			Description: "Cumulative GPA of 90.73/100",
		},
	}

	// Certifications is a slice of Certification struct that contains list of certifications.
	Certifications = []Certification{
		{
			Title:         "Google IT Support Specialization",
			CredentialId:  "EWNAXLA6YJRA",
			CredentialUrl: "https://www.coursera.org/account/accomplishments/specialization/EWNAXLA6YJRA",
		},
		{
			Title:         "AWS Cloud Technical Essentials",
			CredentialId:  "HSKQ9QM64B5A",
			CredentialUrl: "https://www.coursera.org/account/accomplishments/certificate/HSKQ9QM64B5A",
		},
		{
			Title:         "Architecting on AWS (Membangun Arsitektur Cloud di AWS)",
			CredentialId:  "EYX46KOVWPDL",
			CredentialUrl: "https://www.dicoding.com/certificates/EYX46KOVWPDL",
		},
		{
			Title:         "Cloud Practitioner Essentials (Belajar Dasar AWS Cloud)",
			CredentialId:  "JLX1L9V96X72",
			CredentialUrl: "https://www.dicoding.com/certificates/JLX1LVQ32X72",
		},
		{
			Title:         "Menjadi Back-End Developer Expert",
			CredentialId:  "QLZ944E7DP5D",
			CredentialUrl: "https://www.dicoding.com/certificates/QLZ944E7DP5D",
		},
		{
			Title:         "Belajar Fundamental Aplikasi Back-End",
			CredentialId:  "98XWV6M3WPM3",
			CredentialUrl: "https://www.dicoding.com/certificates/98XWV6M3WPM3",
		},
		{
			Title:         "Belajar Membuat Aplikasi Back-End untuk Pemula",
			CredentialId:  "JLX1L9V96X72",
			CredentialUrl: "https://www.dicoding.com/certificates/JLX1LVQ32X72",
		},
		{
			Title:         "Belajar Dasar Pemrograman JavaScript",
			CredentialId:  "JLX1L9V96X72",
			CredentialUrl: "https://www.dicoding.com/certificates/JLX1LVQ32X72",
		},
		{
			Title:         "Belajar Dasar Structured Query Language (SQL)",
			CredentialId:  "2VX3J6603PYQ",
			CredentialUrl: "https://www.dicoding.com/certificates/2VX3J6603PYQ",
		},
		{
			Title:         "Back-End Master Class [Golang + Postgres + Kubernetes + gRPC]",
			CredentialId:  "UC-69e6c3c2-7a66-4744-8b60-c47ea0702897",
			CredentialUrl: "https://www.udemy.com/certificate/UC-69e6c3c2-7a66-4744-8b60-c47ea0702897",
		},
	}

	// Awards is a slice of Award struct that contains list of awards.
	Awards = []Award{
		{
			Place:       1,
			Suffix:      "st",
			Host:        "Brawijaya University",
			Competition: "Lomba Penulisan Proposal Program Kreativitas Mahasiswa (LP2PKM) 2022",
			Translation: "Student Creativity Program Proposal Writing Competition 2022",
		},
		{
			Place:       1,
			Suffix:      "st",
			Host:        "Medan State University",
			Competition: "Business Plan Competition Nasional UNIMED 2021",
			Translation: "UNIMED National Business Plan Competition 2021",
		},
	}

	// SkillsFaIcons is a slice of string that contains list of Font Awesome icons for tech stacks.
	SkillsFaIcons = []string{
		"golang",
		"java",
		"html5",
		"css3-alt",
		"js-square",
		"node-js",
		"php",
		"laravel",
		"npm",
		"git-alt",
		"linux",
		"docker",
		"digital-ocean",
		"aws",
		"google",
	}

	// Workflows is a slice of string that contains list of software development workflows.
	Workflows = []string{"Clean Architecture", "Test-Driven Development", "CI/CD Pipeline", "Agile Methodology"}
)
