function linearFunction(n) {
  let count = 0;
  for (let i = 0; i < n; i++) {
    count++;
  }
  return count;
}

function testPerformance(fn, input) {
  const start = performance.now();
  fn(input);
  const end = performance.now();
  return end - start;
}

const sizes = [100_000, 200_000, 500_000, 1_000_000];

for (let size of sizes) {
  const time = testPerformance(linearFunction, size);
  console.log(`n = ${size}, time = ${time.toFixed(2)} ms`);
}
