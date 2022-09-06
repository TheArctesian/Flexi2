package main

import "fmt"

func main() {
	allRooms := [63]int16{221, 222, 223, 224, 225, 226, 301, 302, 311, 401, 402, 411, 412, 413, 421, 422, 423, 424, 501, 503, 611, 507, 511, 512, 513, 514, 521, 522, 523, 524, 525, 526, 601, 603, 604, 605, 607, 505, 612, 613, 614, 615, 616, 617, 618, 619, 621, 622, 623, 624, 625, 626, 701, 703, 705, 707, 708, 711, 712, 713, 714, 722, 724}
	allTeachers := [63]string(['Yip, Serene', 'Tang, Clara', 'Chan, Anna', 'Adair, Andrea', 'Huang, Yuwei', 'Lui, Caroline', 'Zhou, Cong Faith', 'Niu, Peter', 'Cahusac, Bonnie', 'Munoz, Anne-Mar', 'Sin, Kin Yan', 'Huo, Qian', 'Varro, David', 'Wong, Alison', 'Ward, Janine', 'Halliday, Simon', 'Qian, Qian', 'Vitzikam, Aline', 'Motta, Daniel', 'de Sales, Olivier', 'Melling, Bruna', 'Sutton, Jennifer', 'Silver, Gregory', 'Zhang, Jing', 'Momeni, Sanaz', 'Lagos, Edwin', 'Tomochko, Nathan', "O'Reilly, Josie", 'Ma, Irina', 'Luo, Sha', 'Kwan, Le-Shan', 'Jacobsen, Paul', 'Yeung, Mandy', 'Lin, Eunice', 'Lee, John', "O'Donnell, Jacqueline", 'Chen, Bingmei', 'Stahlberg, Heidi', 'Li, Lynn', 'Murphy, Francis', 'Jiang, Yurou', 'Bacon, Helen', 'Harrington, Alicia', 'Jones, Robert', 'Clouting, Ross', 'Fenton, Ryan', 'Keeman, Katarine', 'Hay, Tiffany', 'Jin, Lifan', 'Chadwick, Jenny', 'Mumm, Andrew C.', 'Sohn, Jennifer', 'King, John', 'Li, Suling', 'Jonny Ishaque', 'Lanyon, David', 'Codrington, Janelle', 'Landete, Alicia', 'Coyle, Anthony', 'Doherty, John', 'Osborne, Nicholas', 'Bullock, Jessica', 'Philipps, Michael'])
	
	fmt.Println(allRooms)
	fmt.Println(allTeachers)
	//create a map of rooms and teachers
	teacherRoomMap := make(map[string][]int16)
	for i := 0; i < len(allRooms); i++ {
		teacherRoomMap[allTeachers[i]] = append(teacherRoomMap[allTeachers[i]], 
										 allRooms[i])
	}
	// post to DB
}
