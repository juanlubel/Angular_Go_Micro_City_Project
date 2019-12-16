export interface Topic {
    Author: String
    TopicTitle: String
  }
export interface Comment {
  Author      :String ,
	Body        :String ,
	TopicTittle :String
}
export interface TopicWithComments {
  Topic:   Topic,
	Comments: Comment[]
}