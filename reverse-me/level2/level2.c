#include <stdio.h>
#include <stdbool.h>
#include <string.h>
#include <stdlib.h>

void no() {
	puts("Nope.");
	exit(1);
}

void ok() {
	puts("Good job.");
}

int main() {
	int int0_48;
	int bool_41; //0x41
	char atoi_buffer[3]; //0x39 - 0x37
	int int0_36; //0x36
	int int0_14; //0x14 
	int int0_10; //0x10 
	char buffer[23]; // 0x35 - 0x1E
	char buffer_2[9]; // 0x1d
	

	printf("Please enter key: ");
	int value = scanf("%23s", buffer);
	if (value != 1) {
		no();
	}
	if (buffer[1] != 0x30) {
		no();
	}
	if (buffer[0] != 0x30) {
		no();
	}
	fflush(stdin);
	memset(buffer_2, 0, 9);
	buffer_2[0] = 0x64;
	int0_36 = 0; 
	int0_14 = 2; 
	int0_10 = 1; 
_main_221:
	while (true) {
		bool_41 = 0;
		int0_48 = int0_14;
		if (strlen(buffer_2) < 0x8) {
			bool_41 = (int0_48 < strlen(buffer));
		}
		if (bool_41 == 0) {
			break ; // jmp main+378
		}
_main_302:
	//  ebp - 0x40
	//  ebp - 0x14	
		atoi_buffer[0] = buffer[int0_14];
		atoi_buffer[1] = buffer[int0_14 + 1];
		atoi_buffer[2] = buffer[int0_14 + 2];

		buffer_2[int0_10] = atoi(atoi_buffer);
		int0_14 = int0_14 + 3;
		int0_10 = int0_10 + 1;
	} // jmp main+221
_main_378:
	buffer_2[int0_10] = 0;
	if (strcmp(buffer_2, "delabere") == 0) {
		ok();
	} else {
		no();
	}
	return 0;
}
