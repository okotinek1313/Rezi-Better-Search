import 'dart:convert'; // Used to decode JSON responses
import 'dart:io'; // Used to read local config file
import 'package:flutter/material.dart'; // Core Flutter UI toolkit
import 'package:http/http.dart' as http; // For making HTTP requests
import 'package:url_launcher/url_launcher.dart'; // To open links in browser

// Root widget for the search page (stateful because UI updates with results)
class SearchPage extends StatefulWidget {
  const SearchPage({super.key});

  @override
  State<SearchPage> createState() => _SearchPageState();
}

// Holds state (data + logic) for SearchPage
class _SearchPageState extends State<SearchPage> {

  // Controls the text input field (lets you read what user typed)
  final _controller = TextEditingController();

  // Stores search results returned from backend
  List<dynamic> _results = [];

  // Reads port from local config file
  int _getPort() {
    final file = File('./resources/shared/config.json'); // Path to config
    final contents = file.readAsStringSync(); // Read file contents
    final config = jsonDecode(contents); // Parse JSON
    return config['port']; // Extract port value
  }

  // Sends search request to backend and updates UI
  void _search() async {
    final res = await http.get(
      Uri.parse("http://localhost:${_getPort()}/search?q=${_controller.text}")
    );

    // Update UI with results (triggers rebuild)
    setState(() => _results = jsonDecode(res.body)['hits']);
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      body: Column(
        children: [

          // Search bar section (pushed down from top)
          Padding(
            padding: const EdgeInsets.only(top: 60),

            child: Center( // centers the search bar horizontally
              child: SizedBox(
                width: 400, // fixed width for search bar

                child: TextField(
                  controller: _controller,

                  decoration: const InputDecoration(
                    hintText: "Search ...", // placeholder text
                    border: OutlineInputBorder(), // adds visible border
                  ),

                  onSubmitted: (_) => _search(),
                ),
              ),
            ),
          ),

          // Results list (fills remaining space)
          Expanded(
            child: ListView.builder(
              itemCount: _results.length,

              itemBuilder: (context, index) {
                final hit = _results[index];

                return Center(
                  // Centers each result horizontally
                  child: SizedBox(
                    width: 600, // controls width of each result card

                    child: Card(
                      // Material card for better visual grouping
                      margin: const EdgeInsets.symmetric(
                        vertical: 8, // space between results
                      ),

                      child: ListTile(
                        // Clickable row inside the card
                        title: Text(
                          hit['title'],
                          textAlign: TextAlign.center, // center title text
                        ),

                        subtitle: Text(
                          hit['site'],
                          textAlign: TextAlign.center, // center subtitle text
                        ),

                        onTap: () async {
                          final url = Uri.parse(hit['link']);

                          if (await canLaunchUrl(url)) {
                            await launchUrl(
                              url,
                              mode: LaunchMode.externalApplication,
                            );
                          }
                        },
                      ),
                    ),
                  ),
                );
              },
            ),
          ),
        ],
      ),
    );
  }
}