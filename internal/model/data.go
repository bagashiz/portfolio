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
			Name:    "X",
			Href:    "https://x.com/bagashiz",
			Rel:     "me",
			FaClass: "x-twitter",
		},
	}

	// Works is a slice of Experience struct that contains list of work experiences.
	Works = []Experience{
		{
			Job:         "Full Stack Engineer (Platform)",
			Url:         "https://moladin.com/",
			Company:     "Moladin",
			Start:       "October 2024",
			End:         "Present",
			Description: "Collaborated in an agile team to develop internal tools, created responsive user interfaces with Next.js and React, and built services with Go that integrate with other microservices while ensuring scalability and seamless operations.",
		},
		{
			Job:         "Back End Developer (Digital Transformation and Productivity)",
			Url:         "https://braincore.id/",
			Company:     "Braincore",
			Start:       "August 2024",
			End:         "October 2024",
			Description: "Developed an internal URL shortener with analytics to improve efficiency, collaborated with team members in an agile environment, and contributed to digital transformation by creating back-end solutions to enhance productivity.",
		},
		{
			Job:         "Junior Back-End Engineer Intern",
			Url:         "https://godentist.co.id/",
			Company:     "GoDentist (PT PERIKSA GIGI INDONESIA)",
			Start:       "February 2024",
			End:         "April 2024",
			Description: "Developed, maintained, and co-designed web service APIs using JavaScript and Node.js, collaborating with cross-functional stakeholders to troubleshoot issues, ensure alignment with business requirements, and maintain optimal performance and user experience.",
		},
		{
			Job:         "Laboratory Teaching Assistant of Database Administration",
			Url:         "https://filkom.ub.ac.id",
			Company:     "Faculty of Computer Science (FILKOM) Brawijaya University",
			Start:       "August 2023",
			End:         "December 2023",
			Description: "Led 10+ hands-on lab sessions for 30+ students, enhancing practical skills in IBM DB2 database administration. Provided guidance on Linux systems and assessed proficiency through well-structured evaluations.",
		},
		{
			Job:         "Laboratory Teaching Assistant of Advanced Programming",
			Url:         "https://filkom.ub.ac.id",
			Company:     "Faculty of Computer Science (FILKOM) Brawijaya University",
			Start:       "February 2023",
			End:         "June 2023",
			Description: "Led 10+ hands-on lab sessions, enhancing Java programming skills for 20+ students. Provided guidance on Java libraries, reinforced OOP concepts, and assessed proficiency through well-structured evaluations.",
		},
	}

	// Volunteers is a slice of Experience struct that contains list of volunteer experiences.
	Volunteers = []Experience{
		{
			Job:         "Vice Chairman",
			Url:         "https://porosfilkom.ub.ac.id",
			Company:     "POROS Organization of Open Source",
			Start:       "February 2024",
			End:         "December 2024",
			Description: "Supported and substituted for the Chairman in conducting organizational tasks, including board meetings, facilitating internal and external communication, and providing guidance to other members.",
		},
		{
			Job:         "Web & Cloud Curriculum Mentor",
			Url:         "https://gdsc.community.dev/university-of-brawijaya",
			Company:     "Google Developer Student Clubs University of Brawijaya Chapter",
			Start:       "October 2023",
			End:         "June 2024",
			Description: "Collaborated with mentors to develop and share web and cloud development materials with 100+ members. Offered mentorship, guidance, and feedback, delivering insights as a workshop speaker to enhance learning experiences.",
		},
		{
			Job:         "DevOps Staff of IT Division",
			Url:         "https://hology.ub.ac.id",
			Company:     "Hology UB 6.0",
			Start:       "May 2023",
			End:         "October 2023",
			Description: "Efficiently packaged and released 2 back-end web API services using GitHub Actions and GitHub Container Registry for seamless CI/CD. Improved scalability and reliability with Docker containers and Nginx web proxy. Ensured uninterrupted web performance through back-end and front-end server maintenance.",
		},
	}

	// Educations is a slice of Education struct that contains list of educational backgrounds.
	Educations = []Education{
		{
			School:      "Brawijaya University",
			Url:         "https://ub.ac.id",
			Major:       "Bachelor of Information Systems",
			Start:       "August 2021",
			End:         "January 2025",
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
		"react",
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
