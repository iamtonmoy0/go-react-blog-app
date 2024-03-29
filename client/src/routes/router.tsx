import { createBrowserRouter } from "react-router-dom";
import App from "../App";
import Blog from "../components/Blog";
import CreateBlog from "../components/CreateBlog";

const router = createBrowserRouter([
  {
    path: "/",
    element: <App />,
    children: [
      { path: "/", element: <Blog /> },
      {
        path: "create",
        element: <CreateBlog />,
      },
    ],
  },
]);

export default router;
