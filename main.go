package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sashabaranov/go-openai"
	"github.com/sashabaranov/go-openai/jsonschema"
)

func main() {
	apiKey := os.Getenv("OPENAI_API_KEY")
	client := openai.NewClient(apiKey)
	ctx := context.Background()

	type Result struct {
		Samples []struct {
			Context  string `json:"context"`
			Question string `json:"question"`
		} `json:"samples"`
	}
	var result Result
	schema, err := jsonschema.GenerateSchemaForType(result)
	if err != nil {
		log.Fatalf("GenerateSchemaForType error: %v", err)
	}
	resp, err := client.CreateChatCompletion(ctx, openai.ChatCompletionRequest{
		Model: openai.GPT4oMini,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "Your are a helpful science tutor for highschool students. Given some content (part of a chapter), create multiple context and question (word range of question should be 5-12, word range of context should be 20-70). Context is section of the content from where you got the question.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Content:  5.1 WHA 5.1 WHA5.1 WHA 5.1 WHAT ARE LIFE PROCESSES? T ARE LIFE PROCESSES? T ARE LIFE PROCESSES? T ARE LIFE PROCESSES? T ARE LIFE PROCESSES? The maintenance functions of living organisms must go on even when they are not doing anything particular. Even when we are just sitting in class, even if we are just asleep, this maintenance job has to go on. The processes which together perform this maintenance job are life processes. Since these maintenance processes are needed to prevent damage and break-down, energy is needed for them. This energy comes from outside the body of the individual organism. So there must be a process to transfer a source of energy from outside the body of the organism, which we call food, to the inside, a process we commonly call nutrition. If the body size of the organisms is to grow, additional raw material will also be needed from outside. Since life on earth depends on carbon- based molecules, most of these food sources are also carbon-based. Depending on the complexity of these carbon sources, different organisms can then use different kinds of nutritional processes. The outside sources of energy could be quite varied, since the environment is not under the control of the individual organism. These sources of energy, therefore, need to be broken down or built up in the body, and must be finally converted to a uniform source of energy that can be used for the various molecular movements needed for maintaining living structures, as well as to the kind of molecules the body needs to grow. For this, a series of chemical reactions in the body are necessary. Oxidising-reducing reactions are some of the most common chemical means to break-down molecules. For this, many organisms use oxygen sourced from outside the body. The process of acquiring oxygen from outside the body, and to use it in the process of break-down of food sources for cellular needs, is what we call respiration. In the case of a single-celled organism, no specific organs for taking in food, exchange of gases or removal of wastes may be needed because the entire surface of the organism is in contact with the environment. But what happens when the body size of the organism increases and the body design becomes more complex? In multi-cellular organisms, all the cells may not be in direct contact with the surrounding environment. Thus, simple diffusion will not meet the requirements of all the cells. We have seen previously how, in multi-cellular organisms, various body parts have specialised in the functions they perform. We are familiar with the idea of these specialised tissues, and with their organisation in the body of the organism. It is therefore not surprising that the uptake of food and of oxygen will also be the function of specialised tissues. However, this poses a problem, since the food and oxygen are now taken up at one place in the body of the organisms, while all parts of the body need them. This situation creates a need for a transportation system for carrying food and oxygen from one place to another in the body. When chemical reactions use the carbon source and the oxygen for energy generation, they create by-products that are not only useless for the cells of the body, but could even be harmful. These waste by- products are therefore needed to be removed from the body and discarded outside by a process called excretion. Again, if the basic rules for body Science design in multi-cellular organisms are followed, a specialised tissue for excretion will be developed, which means that the transportation system will need to transport waste away from cells to this excretory tissue. Let us consider these various processes, so essential to maintain life, one by one.  ",
			},
		},
		ResponseFormat: &openai.ChatCompletionResponseFormat{
			Type: openai.ChatCompletionResponseFormatTypeJSONSchema,
			JSONSchema: &openai.ChatCompletionResponseFormatJSONSchema{
				Name:   "math_reasoning",
				Schema: schema,
				Strict: true,
			},
		},
	})
	if err != nil {
		log.Fatalf("CreateChatCompletion error: %v", err)
	}
	err = schema.Unmarshal(resp.Choices[0].Message.Content, &result)
	if err != nil {
		log.Fatalf("Unmarshal schema error: %v", err)
	}
	for i, item := range result.Samples {

		fmt.Println("\nSample: ", i)
		fmt.Println("Context: ", item.Context)
		fmt.Println("Question: ", item.Question)

	}
}

