package shared

import . "github.com/evertotomalok/text-game/internal/core/domain"

var MainQuestions map[uint]Question = map[uint]Question{

	1: {
		ID:   1,
		Text: "Opss... We have a trouble with our application, and all the system is down. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Tell the team to enter in a war room.",
			RedirectTo:               1,
			DebugComplexityCalculate: -1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Try to debug yorself.",
			RedirectTo:               2,
			DebugComplexityCalculate: 1,
		},
	},
	2: {
		ID:   1,
		Text: "You're writing a new feature, but you found a bug. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Fix the bug.",
			RedirectTo:               3,
			DebugComplexityCalculate: -1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Open a issue.",
			RedirectTo:               4,
			DebugComplexityCalculate: 1,
		},
	},
	3: {
		ID:   1,
		Text: "Ohhhh no! You noticed on slack that something it's wrong with a database's migration. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Change the migration.",
			RedirectTo:               5,
			DebugComplexityCalculate: 1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Do nothing.",
			RedirectTo:               6,
			DebugComplexityCalculate: -1,
		},
	},
	4: {
		ID:   1,
		Text: "Your company's laptop doesn't work and you need to finish a task. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Get your personal laptop.",
			RedirectTo:               7,
			DebugComplexityCalculate: 1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Talk to the infrastructure team.",
			RedirectTo:               8,
			DebugComplexityCalculate: -1,
		},
	},
	5: {
		ID:   5,
		Text: "You are making a review in a pull request, and get some things that could be improved. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Open a change request.",
			RedirectTo:               9,
			DebugComplexityCalculate: -1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Make the changes by yourself.",
			RedirectTo:               10,
			DebugComplexityCalculate: 1,
		},
	},
	6: {
		ID:   6,
		Text: "You are making a planning meet with your team, and you remembered that you discovered a bug during the morning. What will you do?",
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
	7: {
		ID:   7,
		Text: "You are making some queries in production database, and send a delete without where command. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Report it to the SRE team.",
			RedirectTo:               13,
			DebugComplexityCalculate: -1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Try recover it yourself.",
			RedirectTo:               14,
			DebugComplexityCalculate: 1,
		},
	},
	8: {
		ID:   8,
		Text: "You are making some changes in a legacy code, and you need change a piece of code maintained for other team. What will you do?",
		FirstAlternative: QuestionAlternative{
			Text:                     "Ask for help for the other team.",
			RedirectTo:               15,
			DebugComplexityCalculate: -1,
		},
		SecondAlternative: QuestionAlternative{
			Text:                     "Make changes right away, and after you report it to the other team.",
			RedirectTo:               16,
			DebugComplexityCalculate: 1,
		},
	},
}

var ComplementaryQuestions map[uint]ComplementaryQuestion = map[uint]ComplementaryQuestion{
	1: {
		ID:   1,
		Text: "Nice! The team is coming to help you.",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Try to solve all problens only you.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "It's a bad idea.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Share the problems with the team.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Two heads think better than one!",
		},
	},
	2: {
		ID:   2,
		Text: "OOOOhhhh nooooo. Your change didn't help, and the app isn't answer more.",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Try to open a connection with the instance.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Ops! All the environment is down now :(",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Just think in a better solution.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Good choise! Sometime it's better to stop and think better",
		},
	},
	3: {
		ID:   3,
		Text: "Nice try, but it didn't work. What will you do now?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Try a little more.",
			DebugComplexityCalculate: -2,
			TransitionMessage:        "You're a hero. You fix the bug",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Give up.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Okkk. But some time it'll be a problem.",
		},
	},
	4: {
		ID:   4,
		Text: "Cool. The PO registered the task on a board list. What should we do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Get the task to solve?.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "You're so busy, and you can't get more tasks.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Ask to the team if someone doesn't has task, and is available to get it.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Nice done!",
		},
	},
	5: {
		ID:   5,
		Text: "OMG... You've just broken all connections with database, because your migration was rigth, and you changed to a wrong version. And now, who can help you?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Cry to Jesus Christ help you.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "God blessed you, and everything going back normal by itself.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Ask to the SRE team help you.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Ohhh god! All the services are down now because SRE couldn't help you.",
		},
	},
	6: {
		ID:   6,
		Text: "You're a lucky guy. You noticed that migration was right. And now what you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Tell to your friend about this.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Your friend is laughing a lot and making jokes about the fact you almost break the database's migration. You got sad and create a bug in the system",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Do nothing.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Good! Nobody needs to know about that!",
		},
	},
	7: {
		ID:   7,
		Text: "Oh man, you finish the task, but you put some virus on company's network. And now what you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Talk to the SecOps team.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Good. You almost lose your job, but now everything is controlled.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Do nothing.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Oh man! Now the virus has corrupted your files, and you've lost the whole code.",
		},
	},
	8: {
		ID:   8,
		Text: "Your computer is brand new. And now what you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Let's go work.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Good. You have a lot of tasks to do.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Do nothing.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "You forgot that you're delayed, and now you need to do things very fast, and your code has many bugs now.",
		},
	},
	9: {
		ID:   9,
		Text: "Nice, you reported to the PR's owner. And now what you do?",
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
		Text: "Ops, you make some mistakes trying to change the code. And now what you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Revert the code.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Ops, now you have some git problems because you needed to force a push. You created some bugs the code.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Call the PR's owner on slack to show the change.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Nice done. He told you that your logic is wrong, and save the code to receive a new bug.",
		},
	},
	11: {
		ID:   11,
		Text: "Good decision. It was a real bug. And now what you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Report more details.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Good! Now the Product Owner know exactly how to detail bug in a card.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Assing bug fix to you.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "You have so much tasks to do, so you created bugs in other features.",
		},
	},
	12: {
		ID:   12,
		Text: "Unfornutely it was a serious bug, and it has deleted all data from your database. And now what you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Call SRE team to try recover the database.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "You're lucky. The last backup it's been done some minutes before the system stopped.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Cry and go home.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Bad decision. Now you have more one bug to fix.",
		},
	},
	13: {
		ID:   13,
		Text: "SRE team is so busy, and can't help you now. What you will do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Wait more for ask for help again.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Good! Now the SRE team can help you.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Try to fix it yourself.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "OOOOhhhh no! You tried to recover the database, but you almost drop all tables in the database. Now ou have more bugs to fix.",
		},
	},
	14: {
		ID:   14,
		Text: "OOOOhhhh no! You tried to recover the database, but you almost drop all tables in the database. And now what you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Call SRE team to try recover the database.",
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
			TransitionMessage:        "Ohhh no! You broke the contract from some API's, and now you have many bugs to fix.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Wait for the other team complete the task.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Nice! You can back to your primary tasks.",
		},
	},
	16: {
		ID:   16,
		Text: "OOOOhhhh no! You tried to make the changes by yourself, but you create some bugs. And now what you do?",
		FirstAlternative: ComplementaryQuestionAlternative{
			Text:                     "Ask for help to the other team.",
			DebugComplexityCalculate: -1,
			TransitionMessage:        "Good decision! The other team reported you they will change it in the next week.",
		},
		SecondAlternative: ComplementaryQuestionAlternative{
			Text:                     "Try to fix it by yourself.",
			DebugComplexityCalculate: 1,
			TransitionMessage:        "Bad decision. You broke the hole service, and the other team now have many bugs to fix.",
		},
	},
}
