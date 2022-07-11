package shared

import . "github.com/evertotomalok/text-game/internal/core/domain"

var MainQuestions map[uint]Question = map[uint]Question{

	0: {
		ID:   1,
		Text: "Opss... We have a trouble with our application, and all the system is down. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Tell the team to join a war room.",
			RedirectTo:               1,
			DebugComplexityCalculate: -1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Try to debug it yourself.",
			RedirectTo:               2,
			DebugComplexityCalculate: 1,
		},
	},
	1: {
		ID:   2,
		Text: "You're writing a new feature, but you found a bug. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Fix the bug.",
			RedirectTo:               3,
			DebugComplexityCalculate: -1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Open an issue.",
			RedirectTo:               4,
			DebugComplexityCalculate: 1,
		},
	},
	2: {
		ID:   3,
		Text: "Ohhhh no! You noticed on slack that something it's wrong with a database's migration. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Change the database's migration.",
			RedirectTo:               5,
			DebugComplexityCalculate: 1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Do nothing.",
			RedirectTo:               6,
			DebugComplexityCalculate: -1,
		},
	},
	3: {
		ID:   4,
		Text: "Your company's laptop doesn't work, and you must finish a task. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Get your personal laptop, and continue working.",
			RedirectTo:               7,
			DebugComplexityCalculate: 1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Talk to the infrastructure team to fix your company's laptop.",
			RedirectTo:               8,
			DebugComplexityCalculate: -1,
		},
	},
	4: {
		ID:   5,
		Text: "You are making a review in a pull request and noticed some things that could be improved. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Open a change request.",
			RedirectTo:               9,
			DebugComplexityCalculate: -1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Make the changes by yourself and commit.",
			RedirectTo:               10,
			DebugComplexityCalculate: 1,
		},
	},
	5: {
		ID:   6,
		Text: "You are making a planning meeting with your team, and you remembered that you discovered a bug during the morning. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Report it to the team.",
			RedirectTo:               11,
			DebugComplexityCalculate: -1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "You're in doubt if it's a real bug. Do nothing.",
			RedirectTo:               12,
			DebugComplexityCalculate: 1,
		},
	},
	6: {
		ID:   7,
		Text: "You are making some queries in the production database and sending a delete without the where clause. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Report it to the SRE team and your team.",
			RedirectTo:               13,
			DebugComplexityCalculate: -1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Try to recover the data from the database yourself.",
			RedirectTo:               14,
			DebugComplexityCalculate: 1,
		},
	},
	7: {
		ID:   8,
		Text: "You are making some changes in a legacy code, and you need to change a piece of code maintained for another team. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Ask for help from the another team.",
			RedirectTo:               15,
			DebugComplexityCalculate: -1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Make changes right away, and after you report it to another team.",
			RedirectTo:               16,
			DebugComplexityCalculate: 1,
		},
	},
}

var ComplementaryQuestions map[uint]ComplementaryQuestion = map[uint]ComplementaryQuestion{
	1: {
		ID:   1,
		Text: "Nice! The team is coming to help you. What will you do now?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Try to solve all problems only you.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "It's a bad idea. You created some bugs in the system.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Share the problems with the team.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Two heads think better than one! You together could fix the bugs.",
		},
	},
	2: {
		ID:   2,
		Text: "OOOOhhhh no. Your change didn't help, and the app isn't answering more. What will you do now?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Try to open a connection with the instance.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Ops! All the environment is down now :(",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Just think of a better solution.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Good choice! Sometimes it's better to stop and think better.",
		},
	},
	3: {
		ID:   3,
		Text: "Nice try, but it didn't work, and you've just killed some services. What will you do now?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Try a little more.",
			DebugComplexityCalculate: -2,
			TransitionMessage:        "You're a hero. You fixed the bug",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Give up.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Okkk. But some time it'll be a problem, because now you have more bugs in your system.",
		},
	},
	4: {
		ID:   4,
		Text: "Cool. The PO registered the task on a board list. What will you do now?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Get the task to solve.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "You're so busy, and you can't get more tasks. Due to this fact, you created some security failures in your APIs.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Ask the team if someone doesn't have tasks and is available to get them.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Nice done! Now you can return to your tasks.",
		},
	},
	5: {
		ID:   5,
		Text: "OMG... You've just broken all connections with the database because your migration was right, and you changed to the wrong version. And now, who can help you?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Cry to Jesus Christ help you.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "God blessed you, and everything going back to normal by itself.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Ask the SRE team to help you.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Oh god! All the services are down now because SRE couldn't help you.",
		},
	},
	6: {
		ID:   6,
		Text: "You're a lucky guy. You noticed that migration was right. And now what will you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Tell your friend about this.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Your friend is laughing a lot and making jokes about the fact you almost break the database's migration. You got sad and create a bug in the system",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Do nothing.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Awesome! Nobody needs to know about that!",
		},
	},
	7: {
		ID:   7,
		Text: "Oh man, you finish the task, but you put some virus on the company's network. And now what will you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Talk to the SecOps team.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Good. You almost lose your job, but now everything is controlled.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Do nothing.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Oh, buddy! Now the virus has corrupted your files, and you've lost the whole of your code.",
		},
	},
	8: {
		ID:   8,
		Text: "Your computer is brand new. And now what will you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Let's go to work.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Good. You have plenty of tasks to do.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Do nothing.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "You forgot that you're delayed, and now you need to do things very fast, and now your code has many bugs.",
		},
	},
	9: {
		ID:   9,
		Text: "Nice, you reported to the PR's owner. And now what will you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Go to review more PRs.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "It's not a good idea. You're delayed with your tasks.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Continue making your tasks.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Nice done. You've just finished all the things for today.",
		},
	},
	10: {
		ID:   10,
		Text: "Ops, you made some mistakes trying to change the code. And now what will you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Revert the code.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Ops, now you have some git problems because you needed to force a push. You created some bugs in the code.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Call the PR's owner on slack to show the change.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Nice done. He told you that your logic is wrong, and to save the code to receive a new bug.",
		},
	},
	11: {
		ID:   11,
		Text: "Good decision. It was a real bug. And now what will you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Report more details.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Good! Now the Product Owner knows exactly how to detail bugs in a card.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Assign the bug fix to you.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "You have so many tasks to do, so you created bugs in other features.",
		},
	},
	12: {
		ID:   12,
		Text: "Unfortunately, it was a serious bug, and it has deleted all data from your database. And now what will you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Call the SRE team to try to recover the database.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "You're so lucky. The last backup it's been done some minutes before the system stopped.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Just think of a better solution.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "It's not time to think. You must make a decision. Now you have more than one bug to fix.",
		},
	},
	13: {
		ID:   13,
		Text: "SRE team is so busy, and can't help you now. What you will do now?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Wait more for ask for help again.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Good! Now the SRE team can help you.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Try to fix it yourself.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Oh no! You tried to recover the database, but you almost drop all tables in the database. Now you have more bugs to fix.",
		},
	},
	14: {
		ID:   14,
		Text: "Oh no! You tried to recover the database, but you almost drop all tables in the database. And now what will you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Call the SRE team to try to recover the database.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "You're lucky. The last backup it's been done some minutes before the system stopped.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Cry and go home.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Bad decision. Now you have more one bug to fix.",
		},
	},
	15: {
		ID:   15,
		Text: "Good decision! The other team reported you they will change it in the next week. What you will do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Make the change right now, and not wait for the other team.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Ohhh no! You broke the contract from some APIs, and now you have many bugs to fix.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Wait for the other team complete the task.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Nice! You can return to your primary tasks.",
		},
	},
	16: {
		ID:   16,
		Text: "OOOOhhhh no! You tried to make the changes by yourself, but you create some bugs. And now what will you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Ask for help from the other team.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Good decision! The other team reported to you they will change it in the next week.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Try to fix it by yourself.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Bad decision. You broke the hole service, and the other team now has many bugs to fix.",
		},
	},
}
