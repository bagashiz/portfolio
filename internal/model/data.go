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
			End:     "April 2024",
			Descriptions: []string{
				"Developed high-performance web service APIs using JavaScript and Node.js, resulting in a 20% increase in response time.",
				"Advocated for API performance optimization, resulting in a 20% reduction in response times and improving user experience.",
				"Collaborated effectively with cross-functional stakeholders to troubleshoot critical API issues and ensure alignment with business requirements.",
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
				"Managed essential organizational activities as a lead substitute for the Chairman.",
				"Guided and supported other members to enhance collaboration and productivity, resulting in a 20% increase in project completion rate.",
				"Ensured consistent information flow within the organization, serving as a spokesperson for external stakeholders.",
			},
		},
		{
			Job:     "Web & Cloud Curriculum Mentor",
			Url:     "https://gdsc.community.dev/university-of-brawijaya",
			Company: "Google Developer Student Clubs University of Brawijaya Chapter",
			Start:   "October 2023",
			End:     "June 2024",
			Descriptions: []string{
				"Developed and shared learning materials on web and cloud development with more than 100 members, resulting in increased knowledge and skills among the participants.",
				"Guided members in learning web and cloud development through mentorship and feedback, resulting in a 30% increase in project completion rates.",
				"Delivered insights on web and cloud development as a workshop speaker to over 50 members, resulting in a 20% increase in participant engagement.",
			},
		},
		{
			Job:     "DevOps Staff of IT Division",
			Url:     "https://hology.ub.ac.id",
			Company: "Hology UB 6.0",
			Start:   "May 2023",
			End:     "October 2023",
			Descriptions: []string{
				"Managed the packaging and release of 2 back-end web API services through efficient CI/CD pipeline management using GitHub Actions and GitHub Container Registry, resulting in a 20% increase in deployment speed.",
				"Optimized scalability and enhanced reliability of the web API services deployment through the utilization of Docker containers and Nginx web proxy, resulting in a 30% increase in overall system performance.",
				"Managed back-end and front-end server maintenance to ensure uninterrupted performance of the website.",
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
			End:         "June 2024",
			Description: "Distinction Graduate in Cloud Computing Learning Path",
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
			Title:         "Google Cloud Computing Foundations",
			CredentialId:  "23189c34-1653-461c-a1a9-bd4d8837ca10",
			CredentialUrl: "https://www.credly.com/badges/23189c34-1653-461c-a1a9-bd4d8837ca10",
		},
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
			Title:         "Menjadi Google Cloud Engineer",
			CredentialId:  "L4PQ1GOE2XO1",
			CredentialUrl: "https://www.dicoding.com/certificates/L4PQ1GOE2XO1",
		},
		{
			Title:         "Belajar Penerapan Machine Learning dengan Google Cloud",
			CredentialId:  "53XEOYGKKZRN",
			CredentialUrl: "https://www.dicoding.com/certificates/53XEOYGKKZRN",
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
			Title:         "Belajar Membuat Aplikasi Back-End untuk Pemula dengan Google Cloud",
			CredentialId:  "JMZV3G4Q3PN9",
			CredentialUrl: "https://www.dicoding.com/certificates/JMZV3G4Q3PN9",
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
		"python",
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
