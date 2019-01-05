Task
Imagine that you come back from 2 weeks of holidays on a Monday. On the team Kanban board, assigned to you, two tasks await:

As a user, I want to be able to enter my expenses and have them saved for later. As a user, in the application UI, I can navigate to an expenses page. On this page, I can add an expense, setting: The date of the expense The value of the expense The reason of the expense When I click "Save Expense", the expense is then saved in the database. The new expense can then be seen in the list of submitted expenses.
As a user, I want to be able to see a list of my submitted expenses. As a user, in the application UI, I can navigate to an expenses page. On this page, I can see all the expenses I have already submitted in a tabulated list. On this list, I can see: The date of the expense The VAT (Value Added Tax) associated to this expense. VAT is the UK’s sales tax. It is 20% of the value of the expense, and is included in the amount entered by the user. The reason for the expense



run go server

gin run main.go


add a dependecy with dep

dep ensure -add github.com/codegangsta/gin