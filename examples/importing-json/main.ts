import file from "./version.json" assert { type: "json" };
console.log(file.version);
const module = await import("./version.json", {
  assert: { type: "json" },
});
console.log(module.default.version);
