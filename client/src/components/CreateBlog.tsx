import axios from "axios";

export default function CreateBlog() {
  const handleSubmit = (e) => {
    e.preventDefault();
    const form = e.target;
    const title = form.title.value;
    const post = form.description.value;
    axios
      .post("http://localhost:3000/create-blog", { title, post })
      .then((res) => {
        window.alert("blog created", res);
      })
      .catch((err: void) => {
        console.log(err);
      });
  };
  return (
    <div>
      <section className="bg-gray-100">
        <div className="mx-auto max-w-screen-xl px-4 py-16 sm:px-6 lg:px-8">
          <div className="grid grid-cols-1 gap-x-16 gap-y-8 lg:grid-cols-5">
            <div className="rounded-lg bg-white p-8 shadow-lg lg:col-span-3 lg:p-12">
              <form onSubmit={handleSubmit} action="#" className="space-y-4">
                <div>
                  <label className="sr-only" htmlFor="name">
                    Name
                  </label>
                  <input
                    className="w-full rounded-lg border-gray-200 p-3 text-sm"
                    placeholder="Enter Title"
                    type="text"
                    id="title"
                  />
                </div>

                <div>
                  <label className="sr-only" htmlFor="message">
                    write here
                  </label>

                  <textarea
                    className="w-full rounded-lg border-gray-200 p-3 text-sm"
                    placeholder="Message"
                    rows="8"
                    id="description"
                  ></textarea>
                </div>

                <div className="mt-4">
                  <button
                    type="submit"
                    className="inline-block w-full rounded-lg bg-black px-5 py-3 font-medium text-white sm:w-auto"
                  >
                    Post
                  </button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </section>
    </div>
  );
}
