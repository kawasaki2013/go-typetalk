package tests

import (
	"context"
	"testing"
)

func Test_Messages_GetMessage_should_get_a_message(t *testing.T) {
	result, resp, err := client.Messages.GetMessage(context.Background(), topicId, postId)
	test(t, result, resp, err)
}

func Test_Messages_PostMessage_should_post_a_message(t *testing.T) {
	result, resp, err := client.Messages.PostMessage(context.Background(), topicId, "go-typetalk - Test_Messages_PostMessage_should_post_a_message", nil)
	test(t, result, resp, err)
}

func Test_Messages_UpdateMessage_should_update_a_message(t *testing.T) {
	result, resp, err := client.Messages.UpdateMessage(context.Background(), topicId, postId, "go-typetalk - Test_Messages_UpdateMessage_should_update_a_message")
	test(t, result, resp, err)
}

func Test_Messages_GetMessage_should_get_a_message_using_Typetalk_Token(t *testing.T) {
	result, resp, err := clientUsingTypetalkToken.Messages.GetMessage(context.Background(), topicId, postId)
	test(t, result, resp, err)
}

func Test_Messages_PostMessage_should_post_a_message_using_Typetalk_Token(t *testing.T) {
	result, resp, err := clientUsingTypetalkToken.Messages.PostMessage(context.Background(), topicId, "go-typetalk - Test_Messages_PostMessage_should_post_a_message_using_Typetalk_Token", nil)
	test(t, result, resp, err)
}

func Test_Messages_UpdateMessage_should_update_a_message_using_Typetalk_Token(t *testing.T) {
	result, resp, err := clientUsingTypetalkToken.Messages.UpdateMessage(context.Background(), topicId, postId, "go-typetalk - Test_Messages_UpdateMessage_should_update_a_message_using_Typetalk_Token")
	test(t, result, resp, err)
}
