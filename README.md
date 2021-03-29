# matrix
в матрице х*х каждая еденица = случайный цвет, нахожу максимальное кол-во блоков с 1 цветом которые стоят сверху\снизу\лево\права друг от друга


тесты: coverage: 83.3% of statements

benchmark: coverage: 80.0% of statements


BenchmarkCreateMatrix-6                  	13286828	        87.90 ns/op	      32 B/op	       2 allocs/op

BenchmarkSearchInMatrix-6                	 2192186	       552.0 ns/op	     304 B/op	       8 allocs/op

BenchmarkSearchTheBiggestColorsCount-6   	 3554743	       339.2 ns/op	      88 B/op	       3 allocs/op



без оптимизации в функции searchInMatrix о которой написано там в комментарии 

BenchmarkSearchInMatrix-6                	 1243292	       957.5 ns/op	     328 B/op	      16 allocs/op
