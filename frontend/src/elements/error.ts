export function showErrorModal(message: string) {
  const body = document.querySelector("body");
  if (!body) return;

  const div = document.createElement("div");
  div.id = "popup-modal-error";

  div.innerHTML = `
<div class="fixed inset-0 z-50 overflow-auto bg-smoke-light flex justify-center items-center">
  <div class="relative p-4 w-full max-w-md m-auto flex justify-center items-center">
      <div class="relative bg-white rounded-lg shadow dark:bg-gray-700">
          <button id="popup-modal-error-close" type="button" class="absolute top-3 end-2.5 text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ms-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white" data-modal-hide="popup-modal">
              <svg class="w-3 h-3" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 14 14">
                  <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="m1 1 6 6m0 0 6 6M7 7l6-6M7 7l-6 6"/>
              </svg>
              <span class="sr-only">Close modal</span>
          </button>
          <div class="p-4 md:p-5 text-center">
              <svg class="mx-auto mb-4 text-gray-400 w-12 h-12 dark:text-gray-200" aria-hidden="true" xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 20 20">
                  <path stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="2" d="M10 11V6m0 8h.01M19 10a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z"/>
              </svg>
              <h3 class="mb-5 text-lg font-normal text-gray-500 dark:text-gray-400">${message}</h3>
              <button id="popup-modal-error-confirm" data-modal-hide="popup-modal" type="button" class="text-white bg-red-600 hover:bg-red-800 focus:ring-4 focus:outline-none focus:ring-red-300 dark:focus:ring-red-800 font-medium rounded-lg text-sm inline-flex items-center px-5 py-2.5 text-center">
                  Confirm
              </button>
          </div>
      </div>
  </div>
</div>
`;

  body.appendChild(div);

  const close = document.querySelector<HTMLButtonElement>("#popup-modal-error-close");
  const confirm = document.querySelector<HTMLButtonElement>("#popup-modal-error-confirm");

  close?.addEventListener("click", () => {
    const modal = document.querySelector("#popup-modal-error");
    modal?.remove();
  });

  confirm?.addEventListener("click", () => {
    const modal = document.querySelector("#popup-modal-error");
    modal?.remove();
  });
}
