/* eslint-disable @typescript-eslint/no-explicit-any */
export function debounce<F extends (...args: any[]) => any>(
	func: F,
	waitFor: number
): (...args: Parameters<F>) => ReturnType<F> {
	let timeout: ReturnType<typeof setTimeout> | null = null;

	return function (...args: Parameters<F>): ReturnType<F> {
		if (timeout !== null) {
			clearTimeout(timeout);
		}
		timeout = setTimeout(() => func(...args), waitFor);
		return undefined as unknown as ReturnType<F>;
	};
}
