package main

/*
短变量声明 使用说明
*/
func ShortVariableDeclaration() {
	/*
		这里有一个比较微妙的地方：简短变量声明左边的变量可能并不是全部都是刚刚声明的。
		如果有一些已经在相同的词法域声明过了，那么简短变量声明语句对这些已经声明过的变量就只有赋值行为了。
	*/

	// 在下面的代码中，第一个语句声明了in和err两个变量。在第二个语句只声明了out一个变量，然后对已经声明的err进行了赋值操作。
	/*
		in, err := os.Open(infile)
		out, err := os.Create(outfile)

		// 简短变量声明语句中必须至少要声明一个新的变量，下面的代码将不能编译通过：

		f, err := os.Open(infile)
		f, err := os.Create(outfile) // compile error: no new variables
	*/

	// 简短变量声明语句只有对已经在同级词法域声明过的变量才和赋值操作语句等价，
	// 如果变量是在外部词法域声明的，那么简短变量声明语句将会在当前词法域重新声明一个新的变量。
}
