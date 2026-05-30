import { PropsWithChildren } from "react";
import { useLaunch } from "@tarojs/taro";

import "./app.css";

function App({ children }: Readonly<PropsWithChildren>) {
	useLaunch(() => {
		console.log("App launched.");
	});

	return children;
}

export default App;
