import java.util.ArrayList;
import java.util.Collections;
import java.util.List;
import java.util.Scanner;

public class Solution1 {
	public static final int THRESHOLD = 5;

	public static void main(String[] args) {
		System.out.println("Welcome to the Application!");
		System.out.println("Enter 5 valid integers in the range [0, 10]");
		List<Integer> nums = getNums();
		Collections.sort(nums);
		printNumbers(nums);
	}

	private static void printNumbers(List<Integer> nums) {
		for (int num : nums)
			System.out.print(num + " ");
	}

	private static List<Integer> getNums() {
		Scanner scanner = new Scanner(System.in);
		List<Integer> nums = new ArrayList<>();
		while (nums.size() < THRESHOLD) {
			String s = scanner.nextLine();
			if (isNumber(s))
				continue;
			nums.add(Integer.parseInt(s));
		}
		scanner.close();

		return nums;

	}

	private static boolean isNumber(String s) {
		try {
			Integer.parseInt(s);
		} catch (NumberFormatException nfe) {
			System.out.println("Invalid! Try again!");
			return false;
		}
		int num = Integer.parseInt(s);
		if (num < 0 || num > 10) {
			System.out.println("Invalid range! Try again!");
			return false;
		}
		return true;
	}
}
