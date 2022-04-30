package main

import "fmt"
import "piscine"

func output(str string, board [][]string) {
	fmt.Println("pawn_success_maps=========\n");
	for h := 0; h < len(pawn_success_maps); h++{ 
		for i := 0; i < len(pawn_success_maps[h]); i++{ 
			for j := 0; j < len(pawn_success_maps[h][i]); j++{
				fmt.Print(string(pawn_success_maps[h][i][j]))
				fmt.Print(" ")
			}
		fmt.Println("")
	}
		piscine.Checkmate(pawn_success_maps[h])
		fmt.Println("")
	}
}

func main() {
	fmt.Println("valid maps---------\n");

	pawn_success_maps := [][]string{
		{
		"K.",
		".P",
		},
		{
		"K...",
		".P..",
		"....",
		"....",
		},
		{
		"......",
		"......",
		"..K...",
		"...P..",
		"......",
		"......",
		},
		{
		"........",
		"........",
		"........",
		"...K....",
		"....P...",
		"........",
		"........",
		"........",
		},
	}
	fmt.Println("pawn_success_maps=========\n");
	for h := 0; h < len(pawn_success_maps); h++{ 
		for i := 0; i < len(pawn_success_maps[h]); i++{ 
			for j := 0; j < len(pawn_success_maps[h][i]); j++{
				fmt.Print(string(pawn_success_maps[h][i][j]))
				fmt.Print(" ")
			}
		fmt.Println("")
	}
		piscine.Checkmate(pawn_success_maps[h])
		fmt.Println("")
	}

	Bishop_success_maps := [][]string{
		{
		"K.",
		".B",
		},
		{
		"K...",
		".B..",
		"....",
		"....",
		},
		{
		"......",
		"......",
		"..K...",
		"...B..",
		"......",
		"......",
		},
		{
		"........",
		"........",
		"........",
		"...K....",
		"....B...",
		"........",
		"........",
		"........",
		},
	}
	fmt.Println("Bishop_success_maps=========\n");
	for h := 0; h < len(Bishop_success_maps); h++{ 
		for i := 0; i < len(Bishop_success_maps[h]); i++{ 
			for j := 0; j < len(Bishop_success_maps[h][i]); j++{
				fmt.Print(string(Bishop_success_maps[h][i][j]))
				fmt.Print(" ")
			}
		fmt.Println("")
	}
		piscine.Checkmate(Bishop_success_maps[h])
		fmt.Println("")
	}

	queen_success_maps := [][]string{
		{
		"K.",
		".Q",
		},
		{
		"K...",
		".Q..",
		"....",
		"....",
		},
		{
		"......",
		"......",
		"..K...",
		"...Q..",
		"......",
		"......",
		},
		{
		"........",
		"........",
		"........",
		"...K....",
		"....Q...",
		"........",
		"........",
		"........",
		},
	}
	fmt.Println("queen_success_maps=========\n");
	for h := 0; h < len(queen_success_maps); h++{ 
		for i := 0; i < len(queen_success_maps[h]); i++{ 
			for j := 0; j < len(queen_success_maps[h][i]); j++{
				fmt.Print(string(queen_success_maps[h][i][j]))
				fmt.Print(" ")
			}
		fmt.Println("")
	}
		piscine.Checkmate(queen_success_maps[h])
		fmt.Println("")
	}

	success_maps := [][]string{
		{
		"K.",
		".P",
		},
		{
		"K...",
		"....",
		"..B.",
		"....",
		},
		{
		"......",
		"......",
		"..K...",
		"..R...",
		"......",
		"......",
		},
		{
		"........",
		"........",
		"........",
		"...K....",
		"....P...",
		"........",
		"........",
		"........",
		},
	}
	fmt.Println("success_maps=========\n");
	for h := 0; h < len(success_maps); h++{ 
		for i := 0; i < len(success_maps[h]); i++{ 
			for j := 0; j < len(success_maps[h][i]); j++{
				fmt.Print(string(success_maps[h][i][j]))
				fmt.Print(" ")
			}
		fmt.Println("")
	}
		piscine.Checkmate(success_maps[h])
		fmt.Println("")
	}

	fmt.Println("invalid maps---------\n");
	empty_maps := [][]string{
		{
		"..",
		"..",
		},
		{
		"....",
		"....",
		"....",
		"....",
		},
		{
		"......",
		"......",
		"......",
		"......",
		"......",
		"......",
		},
		{
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		"..............................",
		},
    }

	fmt.Println("empty=========\n");
	for h := 0; h < len(empty_maps); h++{ 
		for i := 0; i < len(empty_maps[h]); i++{ 
			for j := 0; j < len(empty_maps[h][i]); j++{
				fmt.Print(string(empty_maps[h][i][j]))
				fmt.Print(" ")
			}
		fmt.Println("")
	}
		piscine.Checkmate(empty_maps[h])
		fmt.Println("")
	}


	wrong_width_maps := [][]string{
		{
		"K",
		"P",
		},
		{
		"K...",
		"....",
		"..B.",
		"...",
		},
		{
		".",
		".",
		"..K...",
		"..R...",
		"......",
		"......",
		},
		{
		"",
		"......",
		"..K...",
		"..R...",
		"......",
		"......",
		},
		{
		"......",
		"",
		"..K...",
		"..R...",
		"...",
		"......",
		},
	}
	fmt.Println("wrong_width_maps=========\n");
	for h := 0; h < len(wrong_width_maps); h++{ 
		for i := 0; i < len(wrong_width_maps[h]); i++{ 
			for j := 0; j < len(wrong_width_maps[h][i]); j++{
				fmt.Print(string(wrong_width_maps[h][i][j]))
				fmt.Print(" ")
			}
		fmt.Println("")
	}
		piscine.Checkmate(wrong_width_maps[h])
		fmt.Println("")
	}

}
