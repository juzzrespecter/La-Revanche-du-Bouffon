#include <stdint.h>
#include <stdbool.h>
#include <string.h>
#include <stdio.h>
#include <stdlib.h>

void ____syscall_malloc() {
	puts("Good job.");
}

void ___syscall_malloc() {
	puts("Nope.");
	exit(1);
}

int main() {
	uint64_t long0_50;
	uint8_t int0_45;
	char atoi_buffer[3]; //0x44, 0x43, 0x42
	uint8_t int0_41; // 0x41
	char buffer[23]; // 0x40
	char memset_buffer[9]; //0x21
	uint8_t int0_18; //0x18
	uint8_t int0_c;
	int int0_10; //0x10
	int int0_8; //0x8

	bool boolean_var; // nunca almacenada en stack

	printf("Please enter key: ");
	int0_8 = scanf("%23s", buffer);
	if (int0_8 != 1) {
		___syscall_malloc();
	}
	if (buffer[1] != '2') {
		___syscall_malloc();
	}
	if (buffer[0] != '4') {
		___syscall_malloc();
	}
	fflush(stdin);
	memset(memset_buffer, 0, 9);
	memset_buffer[0] = '*';
	int0_41 = 0;
	int0_18 = 2;
	int0_c = 1;
_main_167:
	while(true) {	
		int var1 = strlen(memset_buffer);
		int0_45 = 0;
		if (strlen(memset_buffer) < 8) {
			long0_50 = int0_18;
			int0_45 = strlen(buffer);
			boolean_var = long0_50 < int0_45;
		}
_main_227:
		if (boolean_var == 0) {
			break ;
		}
		atoi_buffer[0] = buffer[int0_18];
		atoi_buffer[1] = buffer[int0_18 + 1];
		atoi_buffer[2] = buffer[int0_18 + 2];
		memset_buffer[int0_c] = atoi(atoi_buffer);
		int0_18 += 3;
		int0_c += 1;
	}
_main_321:
	memset_buffer[int0_18] = 0;
	printf("memset buffer: %s\n", memset_buffer);
	int0_10 = strcmp(memset_buffer, "********");
	if (int0_10 == -2)
		___syscall_malloc();
	if (int0_10 == -1)
		___syscall_malloc();
	if (int0_10 == 0) {
		____syscall_malloc();
		goto _main_599;
	}
	if (int0_10 == 1)
		___syscall_malloc();
	if (int0_10 == 2)
		___syscall_malloc();
	if (int0_10 == 3)
		___syscall_malloc();
	if (int0_10 == 4)
		___syscall_malloc();
	if (int0_10 == 5)
		___syscall_malloc();
	if (int0_10 == 0x73)
		___syscall_malloc();
	___syscall_malloc();
_main_599:
	return 0;
}
