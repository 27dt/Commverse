import { useState } from "react";
import "../styles/Chat.css";

function Chat() {
	// State representing current messages as a string array
	const [messages, setMessages] = useState<string[]>([]);

	// Arrow function passed into the form upon submit
	const submitHandler = (formData: FormData) => {
		// Gets the "message" field from the form which we defined and converts to string
		const newMessage: string = formData.get("message") as string;

		// Sets the message state to messages (old) with a spread operator and the new message appended
		setMessages([...messages, newMessage]);
	};

	return (
		<div>
			<div>
				{messages.map((message) => ( // Maps the message state to divs containing that message (arrow function)
					<div>{message}</div>
				))}
			</div>
			<form action={submitHandler}>
				<input name="message" />
				<button type="submit">Send Message</button>
			</form>
		</div>
	);
}

export default Chat;
