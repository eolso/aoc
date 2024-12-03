import 'dart:io';

Map<T, int> countList<T>(List<T> list) {
  Map<T, int> count = {};

  for (var id in list) {
    count.update(
      id,
      (value) => ++value,
      ifAbsent: () => 1,
    );
  }

  return count;
}

(List<int>, List<int>) readFile(String file) {
  List<String> lines = File("input.txt").readAsLinesSync();

  List<int> left = [];
  List<int> right = [];
  for (var line in lines) {
    var splitLine = line.split(RegExp(" +"));

    left.add(int.parse(splitLine[0]));
    right.add(int.parse(splitLine[1]));
  }

  return (left, right);
}
