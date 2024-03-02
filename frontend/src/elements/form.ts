export function showFormModal(name: string, path: string) {
  const body = document.querySelector("body");
  if (!body) return;

  const div = document.createElement("div");
  div.id = "popup-modal-form";

  div.innerHTML = `
<div class="fixed inset-0 z-50 overflow-auto bg-smoke-light flex justify-center items-center">
  <div class="relative p-4 w-full max-w-md max-h-full">
    <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
      <div class="flex items-center justify-between p-4 md:p-5 border-b rounded-t dark:border-gray-600">
        <h3 class="text-lg font-semibold text-gray-900 dark:text-white">Add New Mod</h3>
        <button id="popup-modal-form-close" type="button" class="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-toggle="crud-modal">
          <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
            <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
          </svg>
          <span class="sr-only">Close modal</span>
        </button>
      </div>
      <form class="p-4 md:p-5">
        <div class="grid gap-4 mb-4 grid-cols-2">
          <div class="col-span-2">
            <label for="name" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">Name</label>
            <input type="text" name="name" id="name" value="${name}" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" placeholder="Type mod name" required="">
          </div>
          <div class="col-span-2">
            <label for="file" class="block mb-2 text-sm font-medium text-gray-900 dark:text-white">File</label>
            <input type="text" name="file" id="file" value="${path}" class="bg-gray-50 border border-gray-300 text-gray-900 text-sm rounded-lg focus:ring-primary-600 focus:border-primary-600 block w-full p-2.5 dark:bg-gray-600 dark:border-gray-500 dark:placeholder-gray-400 dark:text-white dark:focus:ring-primary-500 dark:focus:border-primary-500" required="">
          </div>
        </div>
        <button type="submit" class="text-white inline-flex items-center bg-blue-700 hover:bg-blue-800 focus:ring-4 focus:outline-none focus:ring-blue-300 font-medium rounded-lg text-sm px-5 py-2.5 text-center dark:bg-blue-600 dark:hover:bg-blue-700 dark:focus:ring-blue-800">
          <svg class="me-1 -ms-1 w-5 h-5" fill="currentColor" viewBox="0 0 20 20" xmlns="http://www.w3.org/2000/svg"><path fill-rule="evenodd" d="M10 5a1 1 0 011 1v3h3a1 1 0 110 2h-3v3a1 1 0 11-2 0v-3H6a1 1 0 110-2h3V6a1 1 0 011-1z" clip-rule="evenodd"></path></svg>
          Add Mod
        </button>
      </form>
    </div>
  </div>
</div> 
`;

  body.appendChild(div);

  const close = document.querySelector<HTMLButtonElement>("#popup-modal-form-close");
  const form = document.querySelector("form");

  close?.addEventListener("click", () => {
    const modal = document.querySelector("#popup-modal-form");
    modal?.remove();
  });

  form?.addEventListener("submit", (e) => {
    e.preventDefault();

    const name = (form.querySelector("#name") as HTMLInputElement).value;
    const file = (form.querySelector("#file") as HTMLInputElement).value;

    window.addMod(name, file);
    
    const modal = document.querySelector("#popup-modal-form");
    modal?.remove();
  });
}
