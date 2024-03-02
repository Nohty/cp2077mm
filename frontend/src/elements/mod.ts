export function addModElement(container: HTMLElement, name: string) {
  const li = document.createElement("li");
  li.className =
    "flex justify-between items-center bg-gray-700 p-3 rounded my-2 hover:bg-gray-600 transition-colors duration-150";
  li.innerHTML = `
  <span class="text-white font-medium">${name}</span>
  <div class="flex items-center">
    <button onclick="removeMod('${name}')" class="bg-red-600 hover:bg-red-800 text-white font-bold py-1 px-3 rounded text-sm ml-4 transition duration-300">
      Remove
    </button>
  </div>
`;

  container.appendChild(li);
}
